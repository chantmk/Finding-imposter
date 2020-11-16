import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import app from "./app.js";
import random from '../utill/random';
import { Secp256k1Wallet, SigningCosmosClient, makeCosmoshubPath } from "@cosmjs/launchpad";

Vue.use(Vuex);

const API = "http://localhost:1317"
const CHAIN_ID = "Findingimposter"
const ADDRESS_PREFIX = "cosmos"
const LOCAL_STORAGE_LOG_KEY = "finding-imposter-log-secret"
const TYPES = [
  { type: "quarantine", fields: ["user_id", "start_at", "end_at", ] },
  { type: "covid", fields: ["status", "user_id", ] },
  { type: "log", fields: ["place_id", "check_in_at", "check_out_at", ] },
]
const mock = {
  log: [
    {
      id: "123",
      name: "Engineering library",
      checkInAt: "13/6/2020 18:30",
      checkOutAt: null,
    }
  ],
  quarantine: [
    {
      startAt: "13/6/2020 18:30",
      endAt: "13/6/2020 18:30",
    },
  ],
  covid: [
    {
      status: "REJECTED",
      reportAt: "13/6/2020 18:30",
    },
  ]
}

export default new Vuex.Store({
  state: {
    secrets: [],
    data: {
      log: [],
      quarantine: [],
      covid: [],
    }
  },
  mutations: {
    secretsSet(state, { secrets }) {
      state.secrets = secrets;
    },
    secretsUpdate(state, payload) {
      state.secrets.push(payload) 
      localStorage.setItem(LOCAL_STORAGE_LOG_KEY, JSON.stringify(state.secrets));
    },
    dataSet(state, { type, body }) {
      const updated = {};
      updated[type] = body;
      state.data = { ...state.data, ...updated };
    },
  },
  actions: {
    async init({ commit, dispatch }) {
      const _secrets = localStorage.getItem(LOCAL_STORAGE_LOG_KEY);
      console.log(_secrets)
      let secrets
      if(_secrets) secrets = JSON.parse(_secrets)
      else secrets = []
      commit("secretsSet", { secrets });

      // get logs
      await dispatch("getData");
    },
    async getData({ commit }) {
      TYPES.forEach(({ type }) => {
        const body = mock[type]
        commit("dataSet", { type, body });
      });
    },
    async createLog({}, { client, body }) {
      const { data: result } = await axios.post(`${API}/Findingimposter/log`, body);
      const { msg, fee, memo } = result.value;
      // await client.signAndPost(msg, fee, memo);
      const { ID, placeID, createdAt } = msg[0].value
      return { id: ID, placeId: placeID, createdAt }
    },
    async checkout({ dispatch, commit, state }, { logId }) {
      console.log(state.secrets)
      const secret = "churn scrub shrimp course render frost length dinosaur canyon search fog relax belt give drive trouble shove easily"
      try {
        // create new wallet
        const wallet = await Secp256k1Wallet.fromMnemonic(secret, makeCosmoshubPath(0), ADDRESS_PREFIX);
        const { secret: { data }, address } = wallet
        const client = new SigningCosmosClient(API, address, wallet);
        console.log(client, wallet)
        const creator = client.senderAddress

        const body = {
          base_req: { chain_id: CHAIN_ID, from: creator },
          creator,
          logID: logId,
          action: "CHECKOUT"
        }
        const { createdAt } = await dispatch("createLog", { client, body });
        const _log = state.data.log;
        const index = _log.map(i => (i.id)).indexOf(logId)
        _log[index].checkOutAt = createdAt
        commit("dataSet", { type: "log", body: _log });
      } catch(error) {
        console.log(error)
      }

    },
    async checkin({ dispatch, commit, state }, { placeId }) {
      try {
        // create new wallet
        const wallet = await Secp256k1Wallet.generate(18)
        const { secret: { data }, address } = wallet
        const client = new SigningCosmosClient(API, address, wallet);
        const creator = client.senderAddress
        console.log(client, wallet)
        const logID = random()
        const body = {
          base_req: { chain_id: CHAIN_ID, from: creator },
          creator,
          logID,
          placeID: placeId,
          action: "CHECKIN"
        }
        const { id, createdAt } = await dispatch("createLog", { client, body });

        const newLog = { id, name: placeId, placeId, checkInAt: createdAt, checkOutAt: null }
        const _log = state.data.log;
        _log.push(newLog)
        commit("dataSet", { type: "log", body: _log });

        // store secret in local storage
        commit("secretsUpdate", { secret: data, createdAt, logId: logID });
      } catch(error) {
        console.log(error)
      }
    },
    async report({ commit, state }) {
      // create new wallet
      const wallet = await Secp256k1Wallet.generate(18)
      const { secret: { data }, address } = wallet
      
      // create new covid
      const newCovid = { status: "PENDING",  reportAt: "13/6/2020 18:30" }
      const body = state.data.covid;
      body.push(newCovid)
      commit("dataSet", { type: "covid", body });
    },
  },
});
