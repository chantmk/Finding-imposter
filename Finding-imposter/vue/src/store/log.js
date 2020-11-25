import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import random from '../utill/random';
import moment from 'moment';
import { Secp256k1Wallet, SigningCosmosClient, makeCosmoshubPath } from "@cosmjs/launchpad";

Vue.use(Vuex);

const API = "http://localhost:1317"
const CHAIN_ID = "Findingimposter"
const ADDRESS_PREFIX = "cosmos"
const LOCAL_STORAGE_LOG_KEY = "finding-imposter-log-secret"
const LOCAL_STORAGE_COVID_KEY = "finding-imposter-covid-secret"
export default new Vuex.Store({
  state: {
    mnemonic: null,
    secrets: {},
    covidSecrets: {},
    data: {
      log: [],
      quarantine: [],
      covid: [],
    }
  },
  mutations: {
    secretsSet(state, payload) {
      state.secrets = payload;
    },
    secretsUpdate(state, payload) {
      state.secrets = { ...state.secrets, ...payload }
      console.log("secretsUpdate", state.secrets)
      localStorage.setItem(LOCAL_STORAGE_LOG_KEY, JSON.stringify(state.secrets));
    },
    covidSecretsSet(state, payload) {
      state.covidSecrets = payload;
    },
    covidSecretsUpdate(state, payload) {
      state.covidSecrets = { ...state.covidSecrets, ...payload }
      localStorage.setItem(LOCAL_STORAGE_COVID_KEY, JSON.stringify(state.covidSecrets));
    },
    dataSet(state, { type, body }) {
      console.log(type, body)
      const updated = {};
      updated[type] = body? body:[];
      state.data = { ...state.data, ...updated };
    },
    mnemonicSet(state, payload) {
      state.mnemonic = payload;
    }
  },
  actions: {
    async init({ commit, dispatch }) {
      // get local storage : log secret
      // remove 28-day-old address
      const _secrets = localStorage.getItem(LOCAL_STORAGE_LOG_KEY);
      let secrets = JSON.parse(_secrets)
      const expireDate = new moment().subtract(28, "days")
      let filteredSecrets = {}
      for (let key in secrets) {
        const A = new moment(secrets[key].createdAt)
        if(expireDate <= A) filteredSecrets[key] = secrets[key]
      }
      localStorage.setItem(LOCAL_STORAGE_LOG_KEY, JSON.stringify(filteredSecrets));
      commit("secretsSet", filteredSecrets);

      // get local storage : covid secret
      const _covidSecrets = localStorage.getItem(LOCAL_STORAGE_COVID_KEY);
      let covidSecrets = JSON.parse(_covidSecrets)
      commit("covidSecretsSet", covidSecrets);

      // get logs
      dispatch("getLog");
      dispatch("getCovid");
      dispatch("getQuarantine");
    },
    async getLog({ dispatch, commit, state }){
      const { client } = await dispatch("getClient", { isNew: true })
      const creator = client.senderAddress
      const allAddress = Object.values(state.secrets).map(i => i.address)
      const address = allAddress.filter((i, index, self) => (self.indexOf(i) === index))
      const body = {
        base_req: { chain_id: CHAIN_ID, from: creator },
        address,
      }
      const { data: { result } } = await axios.post(`${API}/Findingimposter/log/list`, body);
      let logs = {};
      (result? result:[]).forEach(i => {
        if(!(i.logID in logs)) logs[i.logID] = { logID: i.logID }
        if(i.action == "CHECKOUT") 
          logs[i.logID] = { 
            ...logs[i.logID], 
            checkOutAt: i.createdAt,
            checkOutId: i.id,
          }
        else if(i.action == "CHECKIN") 
          logs[i.logID] = { 
            ...logs[i.logID],
            checkInAt: i.createdAt,
            id: i.id,
            placeID: i.placeID,
            creator: i.creator
          }
      })
      commit("dataSet", { type: "log", body: Object.values(logs) });
    },
    async getCovid({ dispatch, commit, state }){
      const { data: { result } } = await axios.get(`${API}/Findingimposter/covid`);
      // filter
      const ownCovidLog = (result? result:[]).filter(i => i.covidID in state.covidSecrets);
      const covidLog = {};
      ownCovidLog.forEach(i => {
        let data
        if(i.status != "PENDING") {
          data = {
            updatedAt: i.createdAt,
            status: i.status,
            updatedId: i.id,
            doctor: i.creator
          }
        } else {
          const { status, ..._data} = i
          data = _data
        }
        if(i.covidID in covidLog) covidLog[i.covidID] = { status: "PENDING",...covidLog[i.covidID], ...data }
        else covidLog[i.covidID] = { status: "PENDING", ...data }
      })
      await commit("dataSet", { type: "covid", body: Object.values(covidLog) });
    },
    async getQuarantine({ dispatch, commit, state }){
      const { client } = await dispatch("getClient", { isNew: true })
      const creator = client.senderAddress
      const allAddress = Object.values(state.secrets).map(i => i.address)
      const address = allAddress.filter((i, index, self) => (self.indexOf(i) === index))
      const body = {
        base_req: { chain_id: CHAIN_ID, from: creator },
        address,
      }
      const { data: { result } } = await axios.post(`${API}/Findingimposter/quarantine/list`, body);
      // filter
      // const addresses = Object.values(state.secrets).map(i => i.address);
      // const ownQuarantines = result.filter(i => i.userAddress in addresses);
      // await commit("dataSet", { type: "covid", body: ownQuarantines });
      await commit("dataSet", { type: "quarantine", body: result });
    },
    async createLog({}, { client, body }) {
      const { data: result } = await axios.post(`${API}/Findingimposter/log`, body);
      const { msg, fee, memo } = result.value;
      await client.signAndPost(msg, fee, memo);
      const { ID, placeID, createdAt } = msg[0].value
      return { id: ID, placeId: placeID, createdAt }
    },
    async getClient({ state }, { isNew, secret }) {
      const tempSecret = secret? secret : state.mnemonic
      let client;
      let _secret;
      if(false && isNew) {
        const wallet = await Secp256k1Wallet.generate(18)
        const { secret: { data }, address } = wallet;
        _secret = data
        client = new SigningCosmosClient(API, address, wallet);

      } else {
        const wallet = await Secp256k1Wallet.fromMnemonic(tempSecret, makeCosmoshubPath(0), ADDRESS_PREFIX);
        const { secret: { data }, address } = wallet;
        _secret = data
        client = new SigningCosmosClient(API, address, wallet);
      }
      return { client, secret: _secret }
    },
    async checkout({ dispatch, commit, state }, { logId }) {
      try {
        // get client
        const secret = state.secrets[logId].secret;
        const { client } = await dispatch("getClient", { isNew: false, secret })
        const creator = client.senderAddress
        // checkout
        const body = {
          base_req: { chain_id: CHAIN_ID, from: creator },
          creator,
          logID: logId,
          action: "CHECKOUT"
        }
        const { createdAt } = await dispatch("createLog", { client, body });
        // update log
        const _log = state.data.log;
        const index = _log.map(i => (i.logID)).indexOf(logId)
        _log[index].checkOutAt = createdAt
        commit("dataSet", { type: "log", body: _log });
      } catch(error) {
        console.log(error)
      }
    },
    async checkin({ dispatch, commit, state }, { placeId }) {
      try {
        // create new wallet
        const { client, secret } = await dispatch("getClient", { isNew: true })
        const creator = client.senderAddress

        // checkin        
        const logID = random()
        const body = {
          base_req: { chain_id: CHAIN_ID, from: creator },
          creator,
          logID,
          placeID: placeId,
          action: "CHECKIN"
        }
        const { createdAt } = await dispatch("createLog", { client, body });

        // update log
        const newLog = { logID, placeID: placeId, checkInAt: createdAt, checkOutAt: null }
        const _log = state.data.log;
        _log.push(newLog)
        commit("dataSet", { type: "log", body: _log });

        // store secret in local storage
        commit("secretsUpdate", { [logID]: { secret, createdAt, address: creator }});
      } catch(error) {
        console.log(error)
      }
    },
    async report({ dispatch, commit, state }) {
      try {
        // create new wallet
        const { client, secret } = await dispatch("getClient", { isNew: true })
        const creator = client.senderAddress
        const allAddress = Object.values(state.secrets).map(i => i.address)
        const pubKey = allAddress.filter((i, index, self) => (self.indexOf(i) === index))
        // create new covid
        const covidID = random()
        const body = {
          base_req: {
            chain_id: CHAIN_ID,
            from: creator
          },
          creator,
          covidID,
          status: "PENDING",
          pubKey,
        }
        const { data: result } = await axios.post(`${API}/Findingimposter/covid`, body);
        const { msg, fee, memo } = result.value;
        const fee_invalid = {
            amount: [],
            gas: "1000000"
        }
        const a = await client.signAndPost(msg, fee_invalid, memo);
        const { createdAt } = msg[0].value
        if(a.code === 11) throw 'Out of gas';

        // update covid
        const newCovid = { covidID, status: "PENDING",  reportAt: createdAt }
        const newData = state.data.covid;
        newData.push(newCovid)
        commit("dataSet", { type: "covid", body: newData });

        // store secret in local storage
        commit("covidSecretsUpdate", { [covidID]: { secret, address: creator }});
      } catch (error) {
        console.log(error)
      }
    },
    accountSignIn({ commit }, { mnemonic }) {
      commit("mnemonicSet", mnemonic);
    },
  },
});



