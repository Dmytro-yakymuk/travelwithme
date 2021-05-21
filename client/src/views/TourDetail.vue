<template>
    <main v-if="oneTour" id="main">
        <section class="container-fluid trip-info">
            <div class="same-height two-columns row">
                <div class="height col-md-6">
                    <splide v-if="oneTour.images">
                        <splide-slide v-for="image in oneTour.images" :key="image.id" >
                            <img :src="imageURL+image.name" alt="#" height="1104px" width="966px">
                        </splide-slide>
                    </splide>
                </div>
                <div class="height col-md-6 text-col">
                    <div class="holder">
                        <h1 class="small-size">{{oneTour.title}}</h1>
                        <div class="price">
                            від <strong>{{ oneTour.minPrice }} грн</strong>
                        </div>
                        <div class="description">
                            <p>{{oneTour.description}}</p>
                        </div>
                        <ul class="reviews-info">
                            <li>
                                <div class="info-left">
                                    <strong class="title">Рейтинг</strong>
                                    <span class="value">Групи</span>
                                </div>
                                <div class="info-right">
                                    <div class="star-rating">
                                        <span class="value">5/5</span>
                                        <span><span class="icon-star"></span></span>
                                        <span><span class="icon-star"></span></span>
                                        <span><span class="icon-star"></span></span>
                                        <span><span class="icon-star"></span></span>
                                        <span class="disable"><span class="icon-star"></span></span>
                                        
                                    </div>
                                    <span class="value">До {{ oneTour.countGroup }} осіб</span>
                                </div>
                            </li>
                            <li>
                                <div class="info-left">
                                    <strong class="title">Категорія</strong>
                                    <span class="value">Регіон</span>
                                </div>
                                <div class="info-right">
                                    <ul v-if="oneTour.category" class="ico-list">
                                        <span>{{ oneTour.category.name }}</span>
                                        <li>
                                            <span :class="'icon ' + oneTour.category.icon "></span>
                                        </li>
                                        
                                    </ul>
                                    <span v-if="oneTour.region" class="value">{{ oneTour.region.name }} обл.</span>
                                </div>
                            </li>

                        </ul>
                        <div class="btn-holder">
                            <router-link v-scroll-to="'#tab3'" :to="{name: 'dateAndPrice' }" class="btn btn-lg btn-info">Замовити зараз</router-link>
                        </div>
                        
                    </div>
                </div>
            </div>
        </section>
        <div class="tab-container">
            <nav class="nav-wrap" id="sticky-tab">
                <div class="container">
                    <!-- nav tabs -->
                    <ul class="nav nav-tabs text-center">
                        <li>
                            <router-link :to="{name: 'overview'}">Огляд</router-link>
                        </li>
                        <li>
                            <router-link :to="{name: 'comments'}">Відгуки</router-link>
                        </li>
                        <li>
                            <router-link :to="{name: 'dateAndPrice'}">Дати &amp; Ціни</router-link>
                        </li>
                    </ul>
                </div>
            </nav>

            <div class="container tab-content trip-detail">
                 <div class="container">
                    <router-view></router-view>
                </div>
            </div>
            
            <div ref="mapContainer" class="l-map"></div>
        </div>
    </main>
</template>

<script>
    import {mapState, mapGetters, mapMutations, mapActions} from "vuex"
    import {ShowImageURL} from '../constants/index'
    import L from 'leaflet'
    import 'leaflet/dist/leaflet.css'

    // BUG https://github.com/Leaflet/Leaflet/issues/4968
    import iconRetinaUrl from 'leaflet/dist/images/marker-icon-2x.png'
    import iconUrl from 'leaflet/dist/images/marker-icon.png'
    import shadowUrl from 'leaflet/dist/images/marker-shadow.png'

    export default { 
        name: 'TourDetail',
        data: function () {
            return {
                imageURL: ShowImageURL,
            }
        },
        props:['slug'],
        computed: {
            ...mapGetters(['oneTour']),
            ...mapState(['mapInstance']),
        },
        methods: {
            ...mapActions(['getOneTour']),
            ...mapMutations(['setMapInstance']),
            fixBug () {
                // https://github.com/Leaflet/Leaflet/issues/4968
                L.Marker.prototype.options.icon = L.icon({
                    iconRetinaUrl,
                    iconUrl,
                    shadowUrl,
                    iconSize: [25, 41],
                    iconAnchor: [12, 41],
                    popupAnchor: [1, -34],
                    tooltipAnchor: [16, -28],
                    shadowSize: [41, 41]
                })
            },
            createMapInstance () {
                const map = L.map(this.$refs.mapContainer,{ preferCanvas: true }).setView([this.oneTour.x, this.oneTour.y], 15)
                const mapLayer = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
                    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                });
                map.addLayer(mapLayer)

                const marker = L.marker(new L.LatLng(this.oneTour.x, this.oneTour.y))
                map.addLayer(marker)
                return map
            },
            
        },
        async mounted() {
            this.getOneTour({
                "slug": this.slug
            });
            let sleep = ms => new Promise(resolve => setTimeout(resolve, ms));
            await sleep(1000)
            this.setMapInstance(this.createMapInstance())
        },
        beforeDestroy () {
            if (this.mapInstanse) {
                this.mapInstance.remove()
                this.setMapInstance(null)
            }
        },
        created () {
            this.fixBug()
        }
    }
</script>

<style scoped>
    .l-map {
        width: 80%;
        height: 500px;
        margin: 10px 10%;
    }
</style>