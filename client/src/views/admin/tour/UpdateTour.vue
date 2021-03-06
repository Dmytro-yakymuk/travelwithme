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
                                            <input v-model="oneTour.title" class="form-control" placeholder="Введіть назву туру">
                                        </div>

                                        <div class="form-group">
                                            <label>Опис</label>
                                            <textarea v-model="oneTour.description" class="form-control" rows="3"></textarea>
                                        </div>

                                        <div class="form-group">
                                            <label>Категорія</label>
                                            <select v-model="oneTour.category_id" class="form-control">
                                                <option 
                                                    v-for="category in allCategories" :key="category.id"
                                                    :value="category.id"
                                                    :selected="category.id == oneTour.category_id"
                                                >{{category.name}}</option>
                                            </select>
                                        </div>

                                        <div class="form-group">
                                            <label>Регіон</label>
                                            <select v-model="oneTour.region_id"  class="form-control">
                                                <option 
                                                    v-for="region in allRegions" :key="region.id"
                                                    :value="region.id"
                                                    :selected="region.id == oneTour.region_id"
                                                >{{region.name}}</option>
                                            </select>
                                        </div>

                                        <vue-dropzone 
                                            ref="myVueDropzone" 
                                            id="dropzone" 
                                            :options="dropzoneOptions"
                                        ></vue-dropzone>

                                        <ul>
                                            <li v-for="(image, i) in oneTour.images" :key="image.id" class="img_upload" :class="{ 'img_active' : i == 0}">
                                                <img :src="imageURL+image.name" alt="#">
                                                <p @click="deleteImage(image.id)">Видалити</p>
                                            </li>
                                        </ul>
                                        
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
    import {mapGetters, mapActions, mapMutations} from "vuex"

    import vue2Dropzone from 'vue2-dropzone'
    import 'vue2-dropzone/dist/vue2Dropzone.min.css'

    import {AddImageURL, ShowImageURL} from '../../../constants/index'

    export default {
        name: "AdminUpdateTour",
        components: {
            vueDropzone: vue2Dropzone
        },
        data: function () {
            return {
                imageURL: ShowImageURL,
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

        props: ['slug'],
        computed: mapGetters(['oneTour', 'allCategories', 'allRegions']),
        methods: {
            ...mapActions(['getOneTour', 'getAllCategories', 'getAllRegions']),
            ...mapMutations(['updateTourMutation']),

            startProcessingQueue() {
                this.$store.dispatch('updateTour', {
                    'title': this.oneTour.title,
                    'slug': this.oneTour.slug,
                    'description': this.oneTour.description,
                    'status': this.oneTour.status,
                    'category_id': this.oneTour.category_id,
                    'region_id': this.oneTour.region_id
                })
                .then(response => {
                    this.$refs.myVueDropzone.processQueue();
                    this.updateTourMutation(this.getOneTour(this.slug));
                    this.$router.push({ name: 'adminUpdateTour', param: this.slug })
                   
                })
                .catch(error => {
                    if (error.response != null) {
                        if (error.response.status === 401) {
                            this.$router.push({ name: 'logout' })
                        }
                    }
                })
                
            },
            deleteImage(image_id){
                this.$store.dispatch('deleteImage', {
                    'image_id': image_id
                })
                .then(response => {
                    this.updateTourMutation(this.getOneTour(this.slug));
                   // this.$router.push({ name: 'login' })
                   
                })
                .catch(error => {
                    if (error.response != null) {
                        if (error.response.status === 401) {
                            this.$router.push({ name: 'logout' })
                        }
                    }
                })
            }  
        },
        async mounted() {
            this.getOneTour(this.slug);
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