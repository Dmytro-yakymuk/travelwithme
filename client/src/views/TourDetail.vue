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
                            <a href="#" v-scroll-to="'#tab3'" class="btn btn-lg btn-info">Замовити зараз</a>
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
                        <li class="active"><a href="#" v-scroll-to="'#tab1'">Огляд</a></li>
                        <li><a :href="'#tab2'">Відгуки</a></li>
                        <li><a href="#" v-scroll-to="'#tab3'">Дати &amp; Ціни</a></li>
                    </ul>
                </div>
            </nav>
            <!-- tab panes -->
            <div class="container tab-content trip-detail">
                <!-- overview tab content -->
                <div class="tab-pane active" id="tab1">
                    <div class="row">
                        <div class="col-md-6">
                            <strong class="header-box">Детальніше про тур</strong>
                            <div class="detail">
                                <p>This is Photoshop's version  of Lorem Ipsum. Proin gravida nibh vel velit auctor aliquet. Aenean sollicitudin, lorem quis bibendum auctor, nisi elit consequat ipsum, nec sagittis sem nibh id elit. </p>
                                <p>Duis sed odio sit amet nibh vulputate cursus a sit amet mauris. Morbi accumsan ipsum velit. Nam nec tellus a odio tincidunt auctor a ornare odio. </p>
                                <p>Sed non  mauris vitae erat consequat auctor eu in elit. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Mauris in erat justo.</p>
                                <p>Nullam ac urna eu felis dapibus condimentum sit amet a augue. Sed non neque elit. Sed ut imperdiet nisi.</p>
                                <p>Proin condimentum fermentum nunc. Etiam pharetra, erat sed fermentum feugiat, velit mauris egestas quam.</p>
                                <p>Ulins aliquam massa nisl quis neque. Proin condimentum fermentum nunc. Etiam pharetra, erat sed fermentum feugiat, velit mauris egestas quam, ut aliquam massa nisl quis neque. </p>
                                <p>Proin condimentum fermentum nunc. Etiam pharetra, erat sed fermentum feugiat, velit mauris egestas quam.</p>
                                <p>Ulins aliquam massa nisl quis neque. Proin condimentum fermentum nunc. Etiam pharetra, erat sed fermentum feugiat, velit mauris egestas quam, ut aliquam massa nisl quis neque. </p>
                                <p>Ulins aliquam massa nisl quis neque. Proin condimentum fermentum nunc. Etiam pharetra, erat sed fermentum feugiat, velit mauris egestas quam, ut aliquam massa nisl quis neque. </p>
                                <p>Suspendisse gin orci enim.</p>
                            </div>
                        </div>
                        <div class="col-md-6">
                            <strong class="header-box">Огляд включень та виключень з туристичного пакету</strong>
                            <div class="text-box">
                                <div class="holder">
                                    <strong class="title">Що входить у цей тур</strong>
                                    <ul class="content-list tick-list">
                                        <li>All breakfasts, lunches and dinners &amp; dining</li>
                                        <li>All accommodation including tea houses en route</li>
                                        <li>All transportation including taxis and coaches</li>
                                        <li>Flights from Heathrow if booked inc. of flight</li>
                                        <li>Tour and trekking guide for entire journey</li>
                                    </ul>
                                </div>
                            </div>
                            <div class="text-box not-included">
                                <div class="holder">
                                    <strong class="title">Що не входить у цей тур</strong>
                                    <ul class="content-list cross-list">
                                        <li>Travel insurance and other emergencies</li>
                                        <li>Visa fees and entry clearing fees</li>
                                        <li>Single room accommodations</li>
                                        <li>Liquors, beeers and bootled beverages</li>
                                        <li>Photography ccessories like cameras etc.</li>
                                    </ul>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- faq and review tab content -->
                <div class="tab-pane" :id="'tab2'">
                    <div class="row">
                        
                        <div class="col-md-12">
                            <div class="header-box">
                                <a href="#" class="link-right">Залишити відгук</a>
                                <span class="rate-left">
                                    <strong class="title">Загальний рейтинг</strong>
                                    <span class="star-rating">
                                        <span><span class="icon-star"></span></span>
                                        <span><span class="icon-star"></span></span>
                                        <span><span class="icon-star"></span></span>
                                        <span><span class="icon-star"></span></span>
                                        <span class="disable"><span class="icon-star"></span></span>
                                    </span>
                                    <span class="value">4.7/5</span>
                                </span>
                            </div>
                            <div class="comments reviews-body">
                                <div class="comment-holder">
                                    <div class="comment-slot">
                                        <div class="thumb">
                                            <a href="#"><img src="img/thumbs/img-05.jpg" height="50" width="50" alt="image description"></a>
                                        </div>
                                        <div class="text">
                                            <header class="comment-head">
                                                <div class="left">
                                                    <strong class="name">
                                                        <a href="#">Cleona Torez - Spain</a>
                                                    </strong>
                                                    <div class="meta">Reviewed on 14/1/2016</div>
                                                </div>
                                                <div class="right">
                                                    <div class="star-rating">
                                                        <span><span class="icon-star"></span></span>
                                                        <span><span class="icon-star"></span></span>
                                                        <span><span class="icon-star"></span></span>
                                                        <span><span class="icon-star"></span></span>
                                                        <span class="disable"><span class="icon-star"></span></span>
                                                    </div>
                                                    <span class="value">4.7/5</span>
                                                </div>
                                            </header>
                                            <div class="des">
                                                <p>This is Photoshop's version  of Lorem Ipsum. Proin gravida nibh vel velit auctor aliquet. Aenean sollicitudin, lorem quis bibendum auctor, nisi elit consequat ipsum, nec sagittis sem nibh id elit. Duis sed odio sit amet nibh vulputate cursus a sit amet mauris. </p>
                                                <div class="link-holder">
                                                    <a href="#">Read Full Review</a>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="comment-slot">
                                        <div class="thumb">
                                            <a href="#"><img src="img/thumbs/img-06.jpg" height="50" width="50" alt="image description"></a>
                                        </div>
                                        <div class="text">
                                            <header class="comment-head">
                                                <div class="left">
                                                    <strong class="name">
                                                        <a href="#">Cleona Torez - Spain</a>
                                                    </strong>
                                                    <div class="meta">Reviewed on 14/1/2016</div>
                                                </div>
                                                <div class="right">
                                                    <div class="star-rating">
                                                        <span><span class="icon-star"></span></span>
                                                        <span><span class="icon-star"></span></span>
                                                        <span><span class="icon-star"></span></span>
                                                        <span><span class="icon-star"></span></span>
                                                        <span class="disable"><span class="icon-star"></span></span>
                                                    </div>
                                                    <span class="value">4.7/5</span>
                                                </div>
                                            </header>
                                            <div class="des">
                                                <p>This is Photoshop's version  of Lorem Ipsum. Proin gravida nibh vel velit auctor aliquet. Aenean sollicitudin, lorem quis bibendum auctor, nisi elit consequat ipsum, nec sagittis sem nibh id elit. Duis sed odio sit amet nibh vulputate cursus a sit amet mauris. </p>
                                                <div class="link-holder">
                                                    <a href="#">Read Full Review</a>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="comment-slot">
                                        <div class="thumb">
                                            <a href="#"><img src="img/thumbs/img-07.jpg" height="50" width="50" alt="image description"></a>
                                        </div>
                                        <div class="text">
                                            <header class="comment-head">
                                                <div class="left">
                                                    <strong class="name">
                                                        <a href="#">Cleona Torez - Spain</a>
                                                    </strong>
                                                    <div class="meta">Reviewed on 14/1/2016</div>
                                                </div>
                                                <div class="right">
                                                    <div class="star-rating">
                                                        <span><span class="icon-star"></span></span>
                                                        <span><span class="icon-star"></span></span>
                                                        <span><span class="icon-star"></span></span>
                                                        <span><span class="icon-star"></span></span>
                                                        <span class="disable"><span class="icon-star"></span></span>
                                                    </div>
                                                    <span class="value">4.7/5</span>
                                                </div>
                                            </header>
                                            <div class="des">
                                                <p>This is Photoshop's version  of Lorem Ipsum. Proin gravida nibh vel velit auctor aliquet. Aenean sollicitudin, lorem quis bibendum auctor, nisi elit consequat ipsum, nec sagittis sem nibh id elit. Duis sed odio sit amet nibh vulputate cursus a sit amet mauris. </p>
                                                <div class="link-holder">
                                                    <a href="#">Read Full Review</a>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div class="link-more text-center">
                                    <a href="#">Show More Reviews - 75 Reviews</a>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

     
                <!-- dates and prices tab content -->
                <div class="tab-pane" id="tab3">
                    <div class="table-container">
                        <div class="table-responsive">
                            <table class="table table-striped">
                                <thead>
                                    <tr>
                                        <th>
                                            <strong class="date-text">Початок</strong>
                                        </th>
                                        <th>
                                            <strong class="date-text">Завершення</strong>
                                        </th>
                                        <th>
                                            <strong class="date-text">Статус</strong>
                                        </th>
                                        <th>
                                            <strong class="date-text">Ціна</strong>
                                        </th>
                                        <th>
                                            &nbsp;
                                        </th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr>
                                        <td><div class="cell"><div class="middle">Fri 18 Oct '16</div></div></td>
                                        <td><div class="cell"><div class="middle"> Sun 04 Nov '16</div></div></td>
                                        <td><div class="cell"><div class="middle">Available &amp; Guaranteed</div></div></td>
                                        <td><div class="cell"><div class="middle">779 грн</div></div></td>
                                        <td><div class="cell"><div class="middle">
                                            <a href="#" class="btn btn-default">Замовити</a>
                                        </div></div></td>
                                    </tr>
                                    <tr>
                                        <td><div class="cell"><div class="middle">Fri 18 Oct '16</div></div></td>
                                        <td><div class="cell"><div class="middle"> Sun 04 Nov '16</div></div></td>
                                        <td><div class="cell"><div class="middle">Booked &amp; Guaranteed</div></div></td>
                                        <td><div class="cell"><div class="middle">679 грн</div></div></td>
                                        <td><div class="cell"><div class="middle">
                                            <a href="#" class="btn btn-default">Замовити</a>
                                        </div></div></td>
                                    </tr>
                                    
                                    <tr>
                                        <td><div class="cell"><div class="middle">Fri 18 Oct '16</div></div></td>
                                        <td><div class="cell"><div class="middle"> Sun 04 Nov '16</div></div></td>
                                        <td><div class="cell"><div class="middle">Available &amp; Guaranteed</div></div></td>
                                        <td><div class="cell"><div class="middle">779 грн</div></div></td>
                                        <td><div class="cell"><div class="middle">
                                            <a href="#" class="btn btn-default">Замовити</a>
                                        </div></div></td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                        <div class="load-more text-center text-uppercase">
                            <a href="#">Більше</a>
                        </div>
                    </div>
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