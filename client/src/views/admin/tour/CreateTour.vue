<template>
    <div id="page-wrapper" >
        <div id="page-inner">
                <!-- /. ROW  -->
            <div class="row">
                <div class="col-lg-12">
                    <div class="panel panel-default">
                        <div class="panel-body">
                            <div class="row">
                                <div class="col-lg-6">
                                    <form role="form" @submit.prevent="startProcessingQueue">
                                        
                                        <div class="form-group">
                                            <label>Назва туру</label>
                                            <input v-model="title" class="form-control" placeholder="Введіть назву туру">
                                        </div>

                                        <div class="form-group">
                                            <label>Опис</label>
                                            <textarea v-model="description" class="form-control" rows="3"></textarea>
                                        </div>

                                        <div class="form-group">
                                            <label>Категорія</label>
                                            <select v-model="category_id" class="form-control">
                                                <option 
                                                    v-for="category in allCategories" :key="category.id"
                                                    :value="category.id"
                                                    :selected="category.id == category_id"
                                                >{{category.name}}</option>
                                            </select>
                                        </div>

                                        <div class="form-group">
                                            <label>Регіон</label>
                                            <select v-model="region_id"  class="form-control">
                                                <option 
                                                    v-for="region in allRegions" :key="region.id"
                                                    :value="region.id"
                                                    :selected="region.id == region_id"
                                                >{{region.name}}</option>
                                            </select>
                                        </div>

                                        <vue-dropzone 
                                            ref="myVueDropzone" 
                                            id="dropzone" 
                                            :options="dropzoneOptions"
                                        ></vue-dropzone>
                                    
                                        <button type="submit" class="btn btn-default">Зберегти</button>
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

    import vue2Dropzone from 'vue2-dropzone'
    import 'vue2-dropzone/dist/vue2Dropzone.min.css'

    import {AddImageURL} from '../../../constants/index'

    export default {
        name: "AdminCreateTour",
        components: {
            vueDropzone: vue2Dropzone
        },
        data: function () {
            return {
                slug: '',

                title: '',
                description: '',
                status: '',
                category_id: '',
                region_id: '',
                dropzoneOptions: {
                    url: AddImageURL+this.slug,
                    method: 'put',
                    thumbnailWidth: 150,
                    maxFilesize: 0.5,
                    addRemoveLinks: true,
                    autoProcessQueue: false,
                    parallelUploads: 5,
                    uploadMultiple: true,
                    maxFiles: 5,
                    headers: { "Authorization": 'Bearer ' + localStorage.getItem('token')}
                }
            }
        },

        computed: mapGetters(['allCategories', 'allRegions']),
        methods: {
            ...mapActions(['getAllCategories', 'getAllRegions']),

            startProcessingQueue() {
                this.$store.dispatch('createTour', {
                    'title': this.title,
                    'description': this.description,
                    'status': this.status,
                    'category_id': this.category_id,
                    'region_id': this.region_id
                })
                .then(response => {
                    this.slug = response.data.slug;
                    this.$refs.myVueDropzone.dropzone.options.url = AddImageURL+this.slug;
                    this.$refs.myVueDropzone.processQueue();
                    this.$router.push({ name: 'adminGetTours'})
                   
                })
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
            this.getAllCategories();
            this.getAllRegions();
        },
        
    }
</script>

<style scoped>
    .img_upload {
        height: 200px;
        width: 150px;
    }
    .img_active {
        background-color: antiquewhite;
        margin-left: 50px;
    }
</style>