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
                                            <input v-model="title" class="form-control" placeholder="Введіть назву туру">
                                        </div>

                                        <div class="form-group">
                                            <label>Короткий опис</label>
                                            <textarea v-model="description" class="form-control" rows="3"></textarea>
                                        </div>

                                        <div class="form-group">
                                            <label>Детальний опис</label>
                                            <textarea v-model="text" class="form-control" rows="5"></textarea>
                                        </div>

                                        <div class="form-group">
                                            <label>Місце туру</label>
                                            <div id="map" style="height: 500px; z-index: 1" />
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
                                        
                                        <div class="form-group">
                                            <label>Події</label>
                                            <div v-for="event in allEvents" :key="event.id" >
                                                <input type="checkbox" :value="event.id" v-model="selected_events">
                                                <label>{{event.text}}</label><br>
                                            </div>
                                        </div>

                                        <vue-dropzone 
                                            ref="myVueDropzone" 
                                            id="dropzone" 
                                            :options="dropzoneOptions"
                                        ></vue-dropzone>
                                        
                                        <br>
                                        <div v-if="getSuccessMessage != ''" class="alert alert-success" role="alert">{{getSuccessMessage}}</div>
                                        <div v-if="getFailMessage != ''" class="alert alert-danger" role="alert">{{getFailMessage}}</div>
                                    
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
            Map,
            vueDropzone: vue2Dropzone
                
        },
        data: function () {
            return {
                slug: '',

                title: 'test_title',
                description: 'test_description',
                text: 'test_text',
                x: '50.39706009821494',
                y: '30.47297291728477',
                status: '',
                category_id: 1,
                region_id: 1,
                selected_events: [],
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

                map: null,
                markerCounter: 0,
                snackbar: false,
                icon: null,
                coords: [],
                layerGroup: null
            }
        },

        computed: mapGetters(['allCategories', 'allRegions', 'allEvents', 'getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['getAllCategories', 'getAllRegions', 'getAllEvents']),

            startProcessingQueue() {
                
                this.$store.dispatch('createTour', {
                    'title': this.title,
                    'description': this.description,
                    'text': this.text,
                    'x': this.x,
                    'y': this.y,
                    'status': this.status,
                    'category_id': this.category_id,
                    'region_id': this.region_id,
                    'selected_events': this.selected_events
                })
                .then(response => {
                    this.slug = response.data.result.slug;
                    this.$refs.myVueDropzone.dropzone.options.url = AddImageURL+this.slug;
                    this.$refs.myVueDropzone.processQueue();
                    // this.$router.push({ name: 'adminGetTours'})
                })
                .catch(error => {
                    if (error.response != null) {
                        if (error.response.status === 401) {
                            this.$router.push({ name: 'logout' })
                        }
                    }
                })  
            },

            onMapClick (e) {
                if (this.layerGroup) {
                    this.layerGroup.clearLayers();
                }
                this.layerGroup = L.layerGroup().addTo(this.map);
                L.marker(e.latlng, {icon: this.icon}).addTo(this.layerGroup);
                this.x = e.latlng.lat
                this.y = e.latlng.lng
            }
        },
        async mounted() {
            this.getAllCategories();
            this.getAllRegions();
            this.getAllEvents();

            this.map = L.map('map', {
                center: [50.25181858045, 27.75506705161105],
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
            this.map.on('click', this.onMapClick)
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