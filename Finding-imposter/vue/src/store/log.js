import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import random from '../utill/random';
import { Secp256k1Wallet, SigningCosmosClient, makeCosmoshubPath } from "@cosmjs/launchpad";

Vue.use(Vuex);

const getStatus = (a, b) => {
  if(a == "PENDING") return b;
  if(b == "PENDING") return a;
  return a;
}
const API = "https://1317-f4e5278e-f118-47aa-aa60-901ab6359533.ws-us02.gitpod.io";
// const API = "http://localhost:1317"
const CHAIN_ID = "Findingimposter"
const ADDRESS_PREFIX = "cosmos"
const LOCAL_STORAGE_LOG_KEY = "finding-imposter-log-secret"
const LOCAL_STORAGE_COVID_KEY = "finding-imposter-covid-secret"
const TYPES = [
  { type: "quarantine", fields: ["user_id", "start_at", "end_at", ] },
  { type: "covid", fields: ["status", "user_id", ] },
  { type: "log", fields: ["place_id", "check_in_at", "check_out_at", ] },
]
const mock = {
  log: [],
  quarantine: [],
  covid: []
}

export default new Vuex.Store({
  state: {
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
  },
  actions: {
    async init({ commit, dispatch }) {
      // reset local storage
      // remove 14-day-old address
      // localStorage.setItem(LOCAL_STORAGE_LOG_KEY, JSON.stringify({ }));
      // localStorage.setItem(LOCAL_STORAGE_COVID_KEY, JSON.stringify({ }));

      // get local storage : log secret
      const _secrets = localStorage.getItem(LOCAL_STORAGE_LOG_KEY);
      let secrets = JSON.parse(_secrets)
      commit("secretsSet", secrets);

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
      const address = Object.values(state.secrets).map(i => i.address)
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
      console.log('getCovid',result)
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
      const address = Object.values(state.secrets).map(i => i.address)
      const body = {
        base_req: { chain_id: CHAIN_ID, from: creator },
        address,
      }
      const { data: { result } } = await axios.post(`${API}/Findingimposter/quarantine/list`, body);
      console.log("getQuarantine", result)
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
    async getClient({}, { isNew, secret = "pluck much casual country tape praise conduct real mask pipe hospital area where sleep brown clip immune acoustic hospital fringe bundle power caution chalk" }) {
      let client;
      let _secret;
      if(false && isNew) {
        const wallet = await Secp256k1Wallet.generate(18)
        const { secret: { data }, address } = wallet;
        _secret = data
        client = new SigningCosmosClient(API, address, wallet);
        
      } else {
        const wallet = await Secp256k1Wallet.fromMnemonic(secret, makeCosmoshubPath(0), ADDRESS_PREFIX);
        const { secret: { data }, address } = wallet;
        _secret = data
        client = new SigningCosmosClient(API, address, wallet);
      }
      return { client, secret: _secret }
    },
    async checkout({ dispatch, commit, state }, { logId }) {
      try {
        // get client
        console.log(state.secrets)
        console.log(logId, state.secrets[logId])
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
        
        // create new covid
        const covidID = random()
        const body = {
          base_req: {
            chain_id: "Findingimposter",
            from: creator
          },
          creator,
          covidID,
          status: "PENDING",
          pubKey: Object.values(state.secrets).map(i => i.address),
        }
        const { data: result } = await axios.post(`${API}/Findingimposter/covid`, body);
        const { msg, fee, memo } = result.value;
        const a = await client.signAndPost(msg, fee, memo);

        // update covid
        const { createdAt } = msg[0].value
        console.log('report', createdAt,a)
        
        const newCovid = { covidID, status: "PENDING",  reportAt: createdAt }
        const newData = state.data.covid;
        newData.push(newCovid)
        commit("dataSet", { type: "covid", body: newData });

        // store secret in local storage
        commit("covidSecretsUpdate", { [covidID]: { secret, createdAt, address: creator }});
      } catch (error) {
        console.log(error)
      }
    },
  },
});



