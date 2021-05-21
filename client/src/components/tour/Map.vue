<template>
    <div class="tab-pane active" id="tab1">
        <div class="row">
            <div ref="mapContainer" class="l-map"></div>
        </div>
    </div>
</template>

<script>   
    import {mapState, mapGetters, mapMutations} from "vuex"

    import L from 'leaflet'
    import 'leaflet/dist/leaflet.css'

    // BUG https://github.com/Leaflet/Leaflet/issues/4968
    import iconRetinaUrl from 'leaflet/dist/images/marker-icon-2x.png'
    import iconUrl from 'leaflet/dist/images/marker-icon.png'
    import shadowUrl from 'leaflet/dist/images/marker-shadow.png'

    export default {
        name: 'map',
        computed: {
            ...mapGetters(['oneTour']),
            ...mapState(['mapInstance']),
        },   
        methods: {
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