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
      const secrets = JSON.parse(_secrets)
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
    checkout({ commit, state }, { logId }) {
      const _secret = state.secrets.filter(i => i.logId === logId)
      const { secret } = _secret[0]

      // create checkout / connect to API
      const checkoutAt = "13/6/2020 18:30"

      const body = state.data.log;
      const index = body.findIndex(i => i.id === logId)
      body[index].checkOutAt = checkoutAt
      commit("dataSet", { type: "log", body });
    },
    async checkin({ commit, state }, { placeId }) {
      try {
        console.log(placeId)
        // create new wallet
        const wallet = await Secp256k1Wallet.generate(18)
        const { secret: { data }, address } = wallet
        const client = new SigningCosmosClient(API, address, wallet);
        const creator = client.senderAddress
        console.log(client)

        const logID = random()
        const body = {
          base_req: { chain_id: CHAIN_ID, from: creator },
          creator,
          logID,
          placeID: placeId,
          action: "CHECKIN"
        }
        const { data: result } = await axios.post(`${API}/Findingimposter/log`, { body });
        const { ID, placeID, createdAt: checkInAt } = result.value.msg[0].value
        const newLog = { id: ID, name: placeID, placeId: placeID, createdAt: checkInAt, checkOutAt: null }
        const _log = state.data.log;
        _log.push(newLog)
        commit("dataSet", { type: "log", body });

        // store secret in local storage
        commit("secretsUpdate", { secret: data, createdAt: checkInAt}); // do we need to store log id
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
