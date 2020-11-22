import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import app from "./app.js";
import { Secp256k1Wallet, SigningCosmosClient, makeCosmoshubPath } from "@cosmjs/launchpad";

Vue.use(Vuex);

const API = "https://1317-f4e5278e-f118-47aa-aa60-901ab6359533.ws-us02.gitpod.io";
// const API = "http://localhost:1317";
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
    async init({ dispatch, commit, state }) {
      await dispatch("getData")
    },
    // async chainIdFetch({ commit }) {
    //   const node_info = (await axios.get(`${API}/node_info`)).data.node_info;
    //   commit("chainIdSet", { chain_id: node_info.network });
    // },
    async getData({ state, commit }) {
      const { data } = await axios.get(`${API}/Findingimposter/covid/pending`)
      commit("dataSet", { data: data.result })
    },
    removeCovidId({ state, commit }, { index }) {
      const data = state.data;
      data.splice(index, 1);
      commit("dataSet", { data });
    },
    async action({ state, dispatch }, { status, id, index }) {
      const creator = state.client
      const body = {
        base_req: {
          chain_id: "Findingimposter",
          from: creator.senderAddress
        },
        creator: creator.senderAddress,
        covidID: id,
        status,
        pubKey,
      }
      const { data: result } = await axios.post(`${API}/Findingimposter/covid`, body);
      const { msg, fee, memo } = result.value;
      await state.client.signAndPost(msg, fee, memo);
      const pubKey = state.data[index].pubKey
      dispatch("removeCovidId", { index })
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
          console.log(client, wallet, acc)
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