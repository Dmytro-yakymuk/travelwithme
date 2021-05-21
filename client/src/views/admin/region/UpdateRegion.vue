<template>
    <div id="page-wrapper" >
        <div id="page-inner">

            <div v-if="getSuccessMessage != ''" class="alert alert-success" role="alert">{{getSuccessMessage}}</div>
            <div v-if="getFailMessage != ''" class="alert alert-danger" role="alert">{{getFailMessage}}</div>

            <div class="row">
                <div class="col-lg-12">
                    <div class="panel panel-default">
                        <div class="panel-body">
                            <div class="row">
                                <div class="col-lg-6">
                                    <form role="form" @submit.prevent="submitForm">
                                        
                                        <div class="form-group">
                                            <label>Назва</label>
                                            <input type="text" v-model="oneRegion.name" class="form-control">
                                        </div>
  
                                        <button type="submit" class="btn btn-default">Редагувати</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import {mapGetters, mapActions} from "vuex"

    export default {
        name: "AdminCreateRegion",
        props: ['slug'],
        computed: mapGetters(['oneRegion', 'getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['getOneRegionBySlug', 'updateRegion']),

            submitForm() {
                this.updateRegion({
                    'name': this.oneRegion.name,
                    'slug': this.oneRegion.slug
                })
                .then(response => {})
                .catch(error => {
                    if (error.response != null) {
                        if (error.response.status === 401) {
                            this.$router.push({ name: 'logout' })
                        }
                    }
                })  
            },
        },
        async mounted() {
            this.getOneRegionBySlug(this.slug);
        }
    }
</script>

<style scoped></style>