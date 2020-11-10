import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import app from "./app.js";
import { Secp256k1Wallet, SigningCosmosClient, makeCosmoshubPath } from "@cosmjs/launchpad";

Vue.use(Vuex);

const API = "http://localhost:1317";
const ADDRESS_PREFIX = "cosmos"

export default new Vuex.Store({
  state: {
    app,
    account: {},
    chain_id: "",
    data: {},
    client: null,
  },
  mutations: {
    dataSet(state, { data }) {
      state.data = data;
    },
    accountUpdate(state, { account }) {
      state.account = account;
    },
    // chainIdSet(state, { chain_id }) {
    //   state.chain_id = chain_id;
    // },
    // entitySet(state, { type, body }) {
    //   const updated = {};
    //   updated[type] = body;
    //   state.data = { ...state.data, ...updated };
    // },
    clientUpdate(state, { client }) {
      state.client = client;
    },
  },
  actions: {
    async init({ dispatch, state }) {
      // await dispatch("chainIdFetch");
      // get all logs
      // dispatch("getData");
    },
    // async chainIdFetch({ commit }) {
    //   const node_info = (await axios.get(`${API}/node_info`)).data.node_info;
    //   commit("chainIdSet", { chain_id: node_info.network });
    // },
    async getData({ state, commit }) {
      // const { chain_id } = state;
      // const url = `${API}/${chain_id}/${type}`;
      // const body = (await axios.get(url)).data.result;
      const data = [
        {
          id: "t17ExvqLpLLzgSoM",
          createdAt: "13/6/2020 18:30",
        },
        {
          id: "t17ExvqLpLLzgSoM",
          createdAt: "13/6/2020 18:30",
        }
      ]
      commit("dataSet", { data });
    },
    removeCovidId({ state, commit }, { id }) {
      const data = state.data;
      const index = data.findIndex(i => i.id === id)
      data.splice(index, 1);
      commit("dataSet", { data });
    },
    async action({ state, dispatch }, { status, id }) {
      // TODO
      // const { chain_id } = state;
      // const creator = state.client.senderAddress;
      // const base_req = { chain_id, from: creator };
      // const req = { base_req, creator, ...body };
      // const { data } = await axios.post(`${API}/${chain_id}/${type}`, req);
      // const { msg, fee, memo } = data.value;
      // return await state.client.signAndPost(msg, fee, memo);

      dispatch("removeCovidId", { id })
    },
    async accountSignIn({ commit }, { mnemonic }) {
      console.log('fsdfsfsß')
      return new Promise(async (resolve, reject) => {
        const wallet = await Secp256k1Wallet.fromMnemonic(mnemonic, makeCosmoshubPath(0), ADDRESS_PREFIX);
        const [{ address }] = await wallet.getAccounts();
        const url = `${API}/auth/accounts/${address}`;
        const acc = (await axios.get(url)).data;
        if (acc.result.value.address === address) {
          const account = acc.result.value;
          const client = new SigningCosmosClient(API, address, wallet);
          commit("accountUpdate", { account });
          commit("clientUpdate", { client });
          resolve(account);
        } else {
          reject("Account doesn't exist.");
        }
      });
    },
    // async accountUpdate({ state, commit }) {
    //   const url = `${API}/auth/accounts/${state.client.senderAddress}`;
    //   const acc = (await axios.get(url)).data;
    //   const account = acc.result.value;
    //   commit("accountUpdate", { account });
    // },

  },
});
