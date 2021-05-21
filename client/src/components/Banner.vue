<template>
    <!-- main banner -->
    <div class="banner banner-home">
        <!-- revolution slider starts -->
        <div class="rev_slider_wrapper fullscreen-container" >
            <div class="fullscreenabanner">
                <!-- main image for revolution slider -->
                <img src="../assets/img/banner/img-01.jpg" alt="image description">
            </div>
        </div>
        <div class="banner-text">
			<div class="center-text">
				<form class="trip-form" @submit="submitForm">
					<fieldset>
						<div class="holder">
							<label for="adventure">Виберіть категорію</label>
							<div class="select-holder">
								<select class="trip dark" v-model="category_id" id="adventure">
									<option 
										v-for="category in allCategories.slice(0,6)" 
										:key="category.id"
										:value="category.id">
										{{ category.name }}
									</option>
								</select>

							</div>
						</div>
						
						<div class="holder">
							<label for="destination">Виберіть регіон</label>
							<div class="select-holder">
								<select class="trip dark" v-model="region_id" id="destination">
									<option 
										v-for="region in allRegions.slice(0,6)" 
										:key="region.id" 
										:value="region.id">
										{{ region.name }} обл.
									</option>
								</select>
							</div>
						</div>
						<div class="holder">
							<label for="date">Виберіть дату</label>
							<div class="select-holder">
								<select class="trip dark" v-model="date" id="date">
									<option 
										v-for="d in dates" 
										:key="d.id"
										:value="d.start+'+'+d.end">
										{{format_date(d.start, 'MMM')}} - {{format_date(d.end, 'MMM YYYY')}}
									</option>
								</select>
							</div>
						</div>
						
						<div class="holder">
							<input class="btn btn-trip" type="submit" value="Нумо подорожувати!">
						</div>
					</fieldset>
				</form>
			</div>
		</div>
        <div class="feature-block">
            <div class="holder">
                <ul>

                    <li v-for="category in allCategories.slice(0,8)" :key="category.id">
                        <a :href="'/tours?category_id='+category.id">
                            <span class="ico">
                                <span :class="category.icon"></span>
                            </span>
                            <span class="info">{{ category.name }}</span>
                        </a>
                    </li>
                    
                </ul>
            </div>
        </div>
    </div>
</template>

<script>
import {mapGetters, mapActions} from "vuex"
import moment from 'moment'

export default {
    name: "banner",
	data: function () {
		return {
			category_id: '',
			region_id: '',
			date: '',
			dates: [],
		}
	},
    computed: mapGetters(['allCategories', 'allRegions']),
    methods: {
		...mapActions(['getAllCategories', 'getAllRegions']),
		submitForm() {
			let query =  JSON.parse(JSON.stringify({
				category_id: this.category_id || undefined,
				region_id: this.region_id || undefined,
				date: this.date || undefined,
			}))
			this.$router.push({ path: "/tours", 
				query
			})
		},
		
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
    }
}   
</script>

<style scoped>
	.fullscreenabanner {
		height: 820px;
	}
	
	.trip-form {
		padding: 20px 10px 0 10px;
	}

	.trip {
		color:#fff;
		background-color: #9d9d9d;
		width: 250px;
		height: 35px;
		
	}

	.btn.btn-trip {
		height: 35px;
		border-radius: 0%;
	}
	
</style>