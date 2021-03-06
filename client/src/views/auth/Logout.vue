<template>

</template>

<script>
import {mapActions} from "vuex"
export default {
    methods: mapActions(['destroyToken']),
    async mounted() {
        this.destroyToken();
        this.$router.push({ name: 'login' });
    }
}
</script>
