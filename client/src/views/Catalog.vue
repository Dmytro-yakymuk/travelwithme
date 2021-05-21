<template>
    <!-- main container -->
	<main id="main">

		<div class="inner-top">
            <div class="container">
                <h1 class="inner-main-heading">Каталог</h1>
                <!-- breadcrumb -->
                <nav class="breadcrumbs">
                    <ul>
                        <li><a href="/">Головна</a></li>
                        <li><span>Каталог</span></li>
                    </ul>
                </nav>
            </div>
        </div>
		
		<!-- content block -->
		<div class="content-block content-sub">
			<div class="container">
				<div class="filter-option">
					<strong class="result-info">Кількість турів що відповідає Вашому пошуку: {{ getCountTours }}</strong>
					
					<div class="select-holder">
						<div class="filter-slide">
							<div class="select-col">
								<select class="filter-select" v-model="category_id">
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
								<select class="filter-select" v-model="region_id">
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
								<select class="filter-select" v-model="date">
									<option value="">Дата</option>
									<option 
										v-for="d in dates" 
										:key="d.id"
										:value="d.start+'+'+d.end">
										{{format_date(d.start, 'MMM')}} - {{format_date(d.end, 'MMM YYYY')}}
									</option>
								</select>
							</div>

							<div class="select-col">
								<select class="filter-select" v-model="price">
									<option value="">Ціна</option>
									<option v-for="price, index in getPrices" :key="index" :value="price.min + '-' + price.max">
										{{ price.min }}{{index != Object.keys(getPrices).length - 1 ? ' - ' + price.max : '+' }} грн
									</option>
								</select>
							</div>
						</div>
					</div>
					
				</div>
				<!-- list view -->
				<div v-if="allTours.length > 0" class="content-holder list-view">

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
				<div class="content-message" v-if="allTours.length == 0">
					<p>Нажаль, турів не знайдено</p>	
				</div>
				<!-- pagination wrap -->
				<Paginate
					v-if="allTours.length > 0"
					v-model="page"
					:page-count="getPageCount"
					:click-handler="pageChangeHandler"
					:prev-text="'<'"
					:next-text="'>'"
					:container-class="'pagination-wrap'">
				</Paginate>
				<!-- <nav v-if="allTours.length > 0" class="pagination-wrap">
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
				</nav> -->
			</div>
		</div>
		
	</main>
</template>

<script>
	import {mapGetters, mapActions, mapMutations, mapState} from "vuex"
	import {ShowImageURL} from '../constants/index'
	import paginationMixin from '@/mixins/pagination.mixin'
	import moment from 'moment'

	export default { 
		name: 'Catalog',
		data: function () {
			return {
				imageURL: ShowImageURL,
				category_id: '',
				region_id: '',
				price: '',
				date: '',
				dates: [],
			}
		},
		mixins: [paginationMixin],
		computed: {
			...mapGetters(['allTours', 'allCategories', 'allRegions', 'getPrices', 'getPageCount', 'getCountTours']),
			//...mapState(['pageCount']),
		},
		methods: {
			...mapActions(['getAllTours', 'getAllCategories', 'getAllRegions']),
			...mapMutations(['updateToursMutation', 'updatePageCount', 'updateCountTours']),

			format_date(value, format='DD MMM YYYY'){
				if (value) {
					return moment(String(value)).format(format); 
				}
			},
		},
		async mounted() {
			this.getAllCategories();
			this.getAllRegions();

			this.dates = [
				{
					'id': 1,
					'start': '2021-09-01',
					'end': '2021-10-01',
				},
				{
					'id': 2,
					'start': '2021-10-01',
					'end': '2021-11-01',
				},
				{
					'id': 3,
					'start': '2021-11-01',
					'end': '2021-12-01',
				}, 
			]

			var query = this.$route.query 

			// перевіряємо на наявність query в url
			if (query.category_id != null) {
				this.category_id = this.$route.query.category_id
			}
			if (query.region_id != null) {
				this.region_id = this.$route.query.region_id
			}
			if (query.price != null) {
				this.price = this.$route.query.price
			}
			if (query.date != null) {
				this.date = this.$route.query.date
			}

			// чекаємо поки підгрузяться тури
			new Promise((resolve, reject) => {
				this.getAllTours({
					'category_id': this.category_id,
					'region_id': this.region_id,
					'price': this.price,
				})
				.then(response => {
					this.updateCountTours(this.allTours.length)
					// передаємо всі тури на пагінацію
					this.setupPagination(this.allTours)
					resolve(response)
				})
				.catch(error => {
					reject(error)
				})
			})

			
		},
		watch: {
			category_id: function () {
				if (this.category_id != '') {

					// додаємо шлях в url
					this.$router.push({ path: this.$route.path, query: {
						category_id: this.category_id,
						region_id: this.$route.query.region_id,
						date: this.date,
						price: this.$route.query.price,
					} })
					
					// обнуляємо пагінацію
					this.page = 1;

					new Promise((resolve, reject) => {
						this.getAllTours({
							'category_id': this.category_id,
							'region_id': this.region_id,
							'date': this.date,
							'price': this.price,
						})
						.then(response => {
							this.setupPagination(this.allTours)
							resolve(response)
						})
						.catch(error => {
							reject(error)
						})
					})
						
				} else {
					this.$router.push({ path: this.$route.path, query: {
						region_id: this.$route.query.region_id,
						date: this.date,
						price: this.$route.query.price,
					} })

					this.page = 1;

					new Promise((resolve, reject) => {
						this.getAllTours({
							'region_id': this.region_id,
							'date': this.date,
							'price': this.price,
						})
						.then(response => {
							this.setupPagination(this.allTours)
							resolve(response)
						})
						.catch(error => {
							reject(error)
						})
					})
				}
			},
			region_id: function () {
				if (this.region_id != '') {

					this.$router.push({ path: this.$route.path, query: {
						category_id: this.$route.query.category_id,
						region_id: this.region_id,
						date: this.date,
						price: this.$route.query.price,
					} })

					this.page = 1;

					new Promise((resolve, reject) => {
						this.getAllTours({
							'category_id': this.category_id,
							'region_id': this.region_id,
							'date': this.date,
							'price': this.price,
						})
						.then(response => {
							this.setupPagination(this.allTours)
							resolve(response)
						})
						.catch(error => {
							reject(error)
						})
					})
						
				} else { 
					this.$router.push({ path: this.$route.path, query: {
						category_id: this.$route.query.category_id,
						date: this.date,
						price: this.$route.query.price,
					} })

					this.page = 1;

					new Promise((resolve, reject) => {
						this.getAllTours({
							'category_id': this.category_id,
							'date': this.date,
							'price': this.price,
						})
						.then(response => {
							this.setupPagination(this.allTours)
							resolve(response)
						})
						.catch(error => {
							reject(error)
						})
					})
				}
			},
			price: function () {
				if (this.price != '') {

					this.$router.push({ path: this.$route.path, query: {
						category_id: this.$route.query.category_id,
						region_id: this.$route.query.region_id,
						date: this.date,
						price: this.price,
					} })

					this.page = 1;

					new Promise((resolve, reject) => {
						this.getAllTours({
							'category_id': this.category_id,
							'region_id': this.region_id,
							'date': this.date,
							'price': this.price,
						})
						.then(response => {
							this.setupPagination(this.allTours)
							resolve(response)
						})
						.catch(error => {
							reject(error)
						})
					})
						
				} else { 
					this.$router.push({ path: this.$route.path, query: {
						category_id: this.$route.query.category_id,
						region_id: this.$route.query.region_id,
						date: this.date,
					} })

					this.page = 1;

					new Promise((resolve, reject) => {
						this.getAllTours({
							'category_id': this.category_id,
							'region_id': this.region_id,
							'date': this.date,
						})
						.then(response => {
							this.setupPagination(this.allTours)
							resolve(response)
						})
						.catch(error => {
							reject(error)
						})
					})
				}
			},

			date: function () {
				if (this.date != '') {

					this.$router.push({ path: this.$route.path, query: {
						category_id: this.$route.query.category_id,
						region_id: this.$route.query.region_id,
						date: this.date,
						price: this.$route.query.price,
					} })

					this.page = 1;

					new Promise((resolve, reject) => {
						this.getAllTours({
							'category_id': this.category_id,
							'region_id': this.region_id,
							'date': this.date,
							'price': this.price,
						})
						.then(response => {
							this.setupPagination(this.allTours)
							resolve(response)
						})
						.catch(error => {
							reject(error)
						})
					})
						
				} else { 
					this.$router.push({ path: this.$route.path, query: {
						category_id: this.$route.query.category_id,
						region_id: this.$route.query.region_id,
						price: this.$route.query.price,
					} })

					this.page = 1;

					new Promise((resolve, reject) => {
						this.getAllTours({
							'category_id': this.category_id,
							'region_id': this.region_id,
							'price': this.price,
						})
						.then(response => {
							this.setupPagination(this.allTours)
							resolve(response)
						})
						.catch(error => {
							reject(error)
						})
					})
				}
			},

			
		}
	}
	
</script>

<style scoped>
	.content-block.content-sub {
		padding: 0;
	}
    .article {
        height: 200px;
    }
    .article > .thumbnail > .img-wrap{
        height: 200px;
    }
	.filter-select {
		color:#fff;
		background-color: #9d9d9d;
		width: 250px;
		height: 35px;
	}
	.select-holder {
		padding: 20px 0;
	}
	.content-message {
		min-height: 232px;
	}
	.select-col {
		margin: 0 70px;
	}
	.select-col:nth-child(1) {
		margin: 0;
	}
</style>