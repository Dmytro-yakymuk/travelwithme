<template>
  <!-- main container -->
  <main id="main">


    <!-- article list holder -->
    <div class="content-block content-spacing">
      <div class="container">
        <header class="content-heading">
          <h2 class="main-heading">Популярні тури</h2>
          <span class="main-subtitle">У нас є унікальний спосіб задовольнити ваші авантюрні очікування!</span>
          <div class="seperator"></div>
        </header>
        <div class="content-holder">
          <div class="row db-3-col">

            <article v-for="tour in allTours" :key="tour.id" class="col-sm-6 col-md-4 article has-hover-s3">
              <div class="img-wrap">
                <img  
                  v-for="image in tour.images.slice(0,1)"
                  :key="image.id"
                  :src="imageURL+image.name" 
                  height="215" width="370">
                <div class="hover-article">
                  <div class="info-footer">
                    <span class="price">від <span>{{ tour.minPrice }} грн</span></span>
                    <a :href="'/tour/'+tour.slug" class="link-more">Детальніше</a>
                  </div>
                </div>
              </div>
              <h3><a :href="'/tour/'+tour.slug">{{ tour.title }}</a></h3>
              <p>{{ tour.description }}</p>
            </article>

          </div>
        </div>
      </div>
    </div>
    <!-- couter block -->
    <!-- <aside class="count-block">
      <div class="container-fluid">
        <header class="content-heading">
          <span class="main-subtitle">Наші тури охоплюють</span>
          <div class="seperator"></div>
        </header>
        <div class="row">
          <div class="col-xs-6 col-md-3 block-1">
            <div class="holder">
              <span class="icon icon-step"></span>
              <span class="info"><span class="counter">8702</span></span>
              <span class="txt">Тварин</span>
            </div>
          </div>
          <div class="col-xs-6 col-md-3 block-2">
            <div class="holder">
              <span class="icon icon-fish-jumping"></span>
              <span class="info"><span class="counter">378</span></span>
              <span class="txt">Риб</span>
            </div>
          </div>
          <div class="col-xs-6 col-md-3 block-3">
            <div class="holder">
              <span class="icon icon-tree"></span>
              <span class="info"><span class="counter">377</span></span>
              <span class="txt">Рослин</span>
            </div>
          </div>
          <div class="col-xs-6 col-md-3 block-4">
            <div class="holder">
              <span class="icon icon-duration"></span>
              <span class="info"><span class="counter">8973</span></span>
              <span class="txt">Днів з природою</span>
            </div>
          </div>
        </div>
      </div>
    </aside> -->

  </main>
</template>

<script>
import {mapGetters, mapActions} from "vuex"
import {ShowImageURL} from '../constants/index'

export default { 
  name: 'Home',
  data: function () {
    return {
      imageURL: ShowImageURL,
    }
  },
  computed: mapGetters(['allTours']),
  methods: mapActions(['getAllTours']),
  async mounted() {
    this.getAllTours({
      'where': null,
      'order_by': null,
      'limit': 6
    });

    // this.getAllUsers({
    //   'where': "role=owner",
    //   'order_by': null,
    //   'limit': 3
    // });
  }
}
</script>
<style scoped>
  .img-wrap{
    height: 300px;
  }
</style>