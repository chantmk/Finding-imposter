<template>
  <div class="password">
    <input
      type="text"
      v-model="password"
      class="password__input"
      placeholder="Password"
    />
    <div class="button" @click="login">
      LOG IN
    </div>
  </div>
</template>

<style scoped>
.container {
  margin-bottom: 1.5rem;
}
.h1 {
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 1rem;
}
.password {
  margin-top: 0.5rem;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.password__input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #E7E7E7;
  font-size: 0.85rem;
  border-radius: 5px;
  color: #000000;
}
.password__input:focus {
  outline: none;
  border: 1px solid black;
}
.password__input::placeholder {
  color: rgba(0, 0, 0, 0.35);
  font-weight: 500;
}
.button {
  text-align: center;
  font-size: 0.85rem;
  background-color: #000000;
  color: #FFFFFF;
  padding: 8px 16px;
  margin-top: 8px;
  border-radius: 5px;
  cursor: pointer;
}
</style>

<script>
import * as bip39 from "bip39";

export default {
  data() {
    return {
      password: "",
      error: false
    };
  },
  computed: {
    account() {
      return this.$store.state.account;
    },
    address() {
      const { client } = this.$store.state;
      const address = client && client.senderAddress;
      return address;
    },
    mnemonicValid() {
      return bip39.validateMnemonic(this.passwordClean);
    },
    passwordClean() {
      return this.password.trim();
    }
  },
  methods: {
    async login() {
      const mnemonic = this.password.trim();
      if (bip39.validateMnemonic(mnemonic) && !this.error) {
        try {
          const { address } = await this.$store.store.dispatch("accountSignIn", { mnemonic })
          if(address) this.$router.push('/doctor'); 
        } catch (error) {
          console.log(error)
        }
      }
    }
  }
};

</script>