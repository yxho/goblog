<template>
  <a-button type="danger" @click="loginOut">Exit</a-button>
</template>

<script>
export default {
  methods: {
    loginOut() {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>
