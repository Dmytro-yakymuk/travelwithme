<template>
    			<!-- main container -->
			<main id="main">
				
				<!-- content block -->
				<div class="content-block content-sub">
					<div class="container">
						<div class="filter-option">
							<strong class="result-info">Кількість турів що відповідає Вашому пошуку: {{ allTours.length }}</strong>
							<div class="layout-holder">
								
								<div class="select-holder">
									<a href="#" class="btn btn-primary btn-filter"><i class="fa fa-sliders"></i> Filter</a>
									<div class="filter-slide">
										<div class="select-col">
											<select class="filter-select">
												<option value="">Категорія</option>
												<option 
                                                    v-for="category in allCategories" 
													:key="category.id"
                                                    :value="category.id"
                                                    :selected="category.id == category_id"
                                                >{{category.name}}</option>
											</select>
										</div>
										<div class="select-col">
											<select class="filter-select">
												<option value="">Регіон</option>
												<option 
                                                    v-for="region in allRegions" 
													:key="region.id"
                                                    :value="region.id"
                                                    :selected="region.id == region_id"
                                                >{{region.name}}</option>
												
											</select>
										</div>
										<div class="select-col">
											<select class="filter-select">
												<option value="Price Range">Ціна</option>
												<option value="Price Range">1 - 499 грн</option>
												<option value="Price Range">500 - 999 грн</option>
												<option value="Price Range">1000 - 1499 грн</option>
												<option value="Price Range">1500 - 2999 грн</option>
												<option value="Price Range">3000+ грн</option>
											</select>
										</div>
									</div>
								</div>
							</div>
						</div>
						<!-- list view -->
						<div class="content-holder list-view">

							<article v-for="tour in allTours" :key="tour.id" class="article has-hover-s1">
								<div class="thumbnail">
									<div class="img-wrap">
                                        <img  
                                            v-for="image in tour.images.slice(0,1)"
                                            :key="image.id"
                                            :src="imageURL+image.name" 
                                            height="200" width="250">
										
									</div>
									<div class="description">
										<div class="col-left">
											<header class="heading">
												<h3><a :href="'/tour/'+tour.slug" >{{tour.title}}</a></h3>
											</header>
											<p>{{tour.description}}</p>
											<div class="reviews-holder">
												<div class="star-rating">
													<span><span class="icon-star"></span></span>
													<span><span class="icon-star"></span></span>
													<span><span class="icon-star"></span></span>
													<span><span class="icon-star"></span></span>
													<span class="disable"><span class="icon-star"></span></span>
												</div>
												<div class="info-rate">Based on 75 Reviews</div>
											</div>
											
										</div>
										<aside class="info-aside">
											<span class="price">від <span>{{ tour.minPrice }} грн</span></span>
											<a :href="'/tour/'+tour.slug" class="btn btn-default">Детальніше</a>
										</aside>
									</div>
								</div>
							</article>

						</div>
						<!-- pagination wrap -->
						<nav class="pagination-wrap">
							<div class="btn-prev">
								<a href="#" aria-label="Previous">
									<span class="icon-angle-right"></span>
								</a>
							</div>
							<ul class="pagination">
								<li><a href="#">1</a></li>
								<li><a href="#">2</a></li>
								<li><a href="#">3</a></li>
								<li class="active"><a href="#">4</a></li>
								<li><a href="#">5</a></li>
								<li>...</li>
								<li><a href="#">7</a></li>
							</ul>
							<div class="btn-next">
								<a href="#" aria-label="Previous">
									<span class="icon-angle-right"></span>
								</a>
							</div>
						</nav>
					</div>
				</div>
				
			</main>
</template>

<script>
import {mapGetters, mapActions} from "vuex"
import {ShowImageURL} from '../constants/index'

export default { 
  name: 'Catalog',
  data: function () {
    return {
      imageURL: ShowImageURL,
	  category_id: '',
	  region_id: ''
    }
  },
  computed: mapGetters(['allTours', 'allCategories', 'allRegions']),
  methods: mapActions(['getAllTours', 'getAllCategories', 'getAllRegions']),
  async mounted() {
    this.getAllTours({
      'where': null,
      'order_by': null,
      'limit': null
    });
	this.getAllCategories();
    this.getAllRegions();

  }
}
</script>

<style scoped>
    .article {
        height: 200px;
    }
    .article > .thumbnail > .img-wrap{
        height: 200px;
    }
</style>