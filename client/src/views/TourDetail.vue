<template>
    <main id="main">
        <!-- main tour information -->
        <section class="container-fluid trip-info">
            <div class="same-height two-columns row">
                <div class="height col-md-6">
                    <!-- top image slideshow -->
                    <div id="tour-slide">
         
                        <div v-for="image in oneTour.images" :key="image.id" class="slide">
                            <!-- class="bg-stretch" -->
                            <div>
                                <img :src="imageURL+image.name" alt="#" height="1104" width="966">
                            </div>
                        </div>

                    </div>
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
                                    <span class="value">75 Відгуків</span>
                                </div>
                                <div class="info-right">
                                    <div class="star-rating">
                                        <span><span class="icon-star"></span></span>
                                        <span><span class="icon-star"></span></span>
                                        <span><span class="icon-star"></span></span>
                                        <span><span class="icon-star"></span></span>
                                        <span class="disable"><span class="icon-star"></span></span>
                                    </div>
                                    <span class="value">5/5</span>
                                </div>
                            </li>
                            <li>
                                <div class="info-left">
                                    <strong class="title">Стиль туру</strong>
                                    <span class="value">Групи </span>
                                </div>
                                <div class="info-right">
                                    <ul class="ico-list">
                                        <li>
                                            <span class="icon icon-hiking"></span>
                                        </li>
                                        <li>
                                            <span class="icon icon-mount"></span>
                                        </li>
                                        <li>
                                            <span class="icon icon-camping"></span>
                                        </li>
                                    </ul>
                                    <span class="value">До 5 осіб</span>
                                </div>
                            </li>

                        </ul>
                        <div class="btn-holder">
                            <router-link v-scroll-to="'#tab3'"  :to="{name: 'dateAndPrice', params: {trips: oneTour.trips} }" class="btn btn-lg btn-info">Замовити зараз</router-link>
                            
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
                            <router-link :to="{name: 'dateAndPrice', params: {trips: oneTour.trips} }">Дати &amp; Ціни</router-link>
                        </li>
                    </ul>
                </div>
            </nav>

            <div class="container tab-content trip-detail">
                <!-- tab content -->
                 <div class="container">
                    <router-view></router-view>
                </div>
            </div>

        </div>
    </main>
</template>

<script>
import {mapGetters, mapActions} from "vuex"
import {ShowImageURL} from '../constants/index'

export default { 
  name: 'TourDetail',
  data: function () {
    return {
      imageURL: ShowImageURL,
    }
  },
  props:['slug'],
  computed: mapGetters(['oneTour']),
  methods: mapActions(['getOneTour']),
  async mounted() {
    this.getOneTour(this.slug);
  }
}
</script>

<style scoped>

</style>