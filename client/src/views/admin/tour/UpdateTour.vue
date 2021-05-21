<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            
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
                                            <label>Короткий опис</label>
                                            <textarea v-model="oneTour.description" class="form-control" rows="3"></textarea>
                                        </div>

                                        <div class="form-group">
                                            <label>Детальний опис</label>
                                            <textarea v-model="oneTour.text" class="form-control" rows="5"></textarea>
                                        </div>

                                        <div class="form-group">
                                            <label>Місце туру</label>
                                            <div id="map" style="height: 500px; z-index: 1" />
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

                                        <div class="form-group">
                                            <label>Події</label>
                                            <div v-for="event in allEvents" :key="event.id" >
                                                <input type="checkbox" :value="event.id" v-model="oneTour.selected_events">
                                                <label>{{event.text}}</label><br>
                                            </div>
                                        </div>

                                        <vue-dropzone 
                                            ref="myVueDropzone" 
                                            id="dropzone" 
                                            :options="dropzoneOptions"
                                        ></vue-dropzone>

                                        <ul class="ul-img">
                                            <li v-for="(image, i) in oneTour.images" :key="image.id" class="img_upload" :class="{ 'img_active' : i == 0}">
                                                <img :src="imageURL+image.name" alt="#">
                                                <a class="btn btn-primary" v-if="role == 'owner'" @click="deleteImage(image.id)">Видалити</a>
                                            </li>
                                        </ul>

                                        <br>
                                        <div v-if="getSuccessMessage != ''" class="alert alert-success" role="alert">{{getSuccessMessage}}</div>
                                        <div v-if="getFailMessage != ''" class="alert alert-danger" role="alert">{{getFailMessage}}</div>
                                        
                                        <button v-if="role == 'owner'" type="submit" class="btn btn-default">Редагувати</button>

                                    </form>

                                    <br>
                                    

                                    <div v-if="allAudits.length > 0" >
                                        <label v-if="allAudits.length > 0">Вказівки</label>

                                        <a v-for="audit, i in allAudits" :key="audit.id" class="list-group-item">
                                            <span class="badge">
                                                <a v-on:click="allAudits.splice(i, 1)" @click="deleteAudit(audit.id)"> <i class="fa fa-trash"></i></a>
                                            </span>
                                            <i class="fa fa-file-text-o"></i> {{ audit.message }}
                                        </a>
                                    </div>



                                    <form v-if="role == 'admin'" role="form" @submit.prevent="sendAudit">
                                        <div class="form-group">
                                            <label>Написати вказівку</label>
                                            <textarea v-model="message" class="form-control" rows="3"></textarea>
                                        </div>

                                        <button type="submit" class="btn btn-default">Надіслати</button>
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
                role: localStorage.getItem('role'),
                dropzoneOptions: {
                    url: AddImageURL+this.slug,
                    method: 'put',
                    thumbnailWidth: 150,
                    maxFilesize: 5,
                    addRemoveLinks: true,
                    autoProcessQueue: false,
                    parallelUploads: 5,
                    uploadMultiple: true,
                    maxFiles: 5,
                    headers: { "Authorization": 'Bearer ' + localStorage.getItem('token')}
                },
                message: '',

                map: null,
                markerCounter: 0,
                snackbar: false,
                icon: null,
                coords: [],
                layerGroup: null
            }
        },

        props: ['slug'],
        computed: mapGetters(['oneTour', 'allCategories', 'allRegions', 'allAudits', 'getSuccessMessage', 'getFailMessage', 'allEvents']),
        methods: {
            ...mapActions(['getOneTour', 'getAllCategories', 'getAllRegions', 'getAuditsByTourId', 'addAudit', 'delAudit', 'getAllEvents']),
            ...mapMutations(['updateTourMutation']),

            startProcessingQueue() {
                this.$store.dispatch('updateTour', {
                    'title': this.oneTour.title,
                    'slug': this.oneTour.slug,
                    'description': this.oneTour.description,
                    'text': this.oneTour.text,
                    'x': this.oneTour.x,
                    'y': this.oneTour.y,
                    'status': this.oneTour.status,
                    'category_id': this.oneTour.category_id,
                    'region_id': this.oneTour.region_id,
                    'selected_events': this.oneTour.selected_events
                })
                .then(response => {
                    this.$refs.myVueDropzone.processQueue()
                    setTimeout(() => {
                        this.updateTourMutation(this.getOneTour({
                            'slug': this.slug
                        }));
                    }, 500)
                })
                .catch(error => {   
                })
                
            },
            sendAudit() {
                this.addAudit({
                    'tour_id': this.oneTour.id,
                    'message': this.message
                })
                
            },
            deleteImage(image_id){
                this.$store.dispatch('deleteImage', {
                    'image_id': image_id
                })
                .then(response => {
                    this.updateTourMutation(this.getOneTour({
                        'slug': this.slug
                    }));
                })
                .catch(error => {
                    if (error.response != null) {
                        if (error.response.status === 401) {
                            this.$router.push({ name: 'logout' })
                        }
                    }
                })
            },
            deleteAudit(auditId) {
                this.delAudit(auditId)
            },

            onMapClick (e) {
                if (this.layerGroup) {
                    this.layerGroup.clearLayers();
                }
                this.layerGroup = L.layerGroup().addTo(this.map);
                L.marker(e.latlng, {icon: this.icon}).addTo(this.layerGroup);
                this.oneTour.x = e.latlng.lat
                this.oneTour.y = e.latlng.lng
            }
        },
        mounted() {
            
            this.getAllCategories();
            this.getAllRegions();
            this.getAuditsByTourId(this.slug);
            this.getAllEvents();

            this.getOneTour({
                'slug': this.slug
            })
            .then(response => {  
                this.map = L.map('map', {
                    center: [this.oneTour.x, this.oneTour.y],
                    zoom: 7
                })
                let leafIcon = L.Icon.extend({
                    options: {
                    iconSize:     [38, 95],
                    shadowSize:   [50, 64],
                    iconAnchor:   [22, 94],
                    shadowAnchor: [4, 62],
                    popupAnchor:  [-3, -76]
                    }
                });

                // маркер - дерево
                this.icon = new leafIcon({
                    iconUrl: 'http://leafletjs.com/examples/custom-icons/leaf-green.png',
                    shadowUrl: 'http://leafletjs.com/examples/custom-icons/leaf-shadow.png'
                })

                // шар для карти TOPO-OSM-WMS
                L.tileLayer.wms('http://ows.mundialis.de/services/service?', {
                    layers: 'TOPO-OSM-WMS',
                    format: 'image/png',
                    transparent: true,
                    attribution: 'Image tiles: &copy <a href="https://land.gov.ua/">ЦДЗК</a>'
                }).addTo(this.map);

                var latlng = L.latLng(this.oneTour.x, this.oneTour.y);
                this.layerGroup = L.layerGroup().addTo(this.map);
                L.marker(latlng, {icon: this.icon}).addTo(this.layerGroup);

                this.map.on('click', this.onMapClick)
            });
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

    ul.ul-img {
        margin: 10px 10px;
        padding: 0;
    }
    ul.ul-img li {
        list-style-type: none;
        margin: 0;
        height: 250px!important;
        width: 200px!important;
    }
    ul.ul-img li img{
        height: 180px;;
    }
   
</style>