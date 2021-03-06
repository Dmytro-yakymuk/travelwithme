<template>
  
</template>

<script>
    
    import {mapActions} from "vuex"
    export default {
        
        name: "AdminDeleteTour",
        props: ['slug'],
        methods: mapActions(['deleteTour']),
        async mounted() {
            this.deleteTour(this.slug);
            this.$router.push({ name: 'adminGetTours' });
        },
        
    }
</script>

<style scoped>

</style>