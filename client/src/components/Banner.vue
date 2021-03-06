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
						<form class="trip-form">
							<fieldset>
								<div class="holder">
									<label for="adventure">Select Your Adventure</label>
									<div class="select-holder">
										<select class="trip dark" id="adventure">
											<option>Jungle Safari</option>
											<option>Wildlife Photography</option>
											<option>Bird Watching</option>
											<option>4 Wheel Drives</option>
											<option>Hunting Fishing</option>
											<option>Discover Africa</option>
										</select>
									</div>
								</div>
								<div class="holder">
									<label for="destination">Select Your Destination</label>
									<div class="select-holder">
										<select class="trip dark" id="destination">
											<option>Europe - Spain</option>
											<option>Europe - Greece</option>
											<option>Africa - Kenya</option>
											<option>Africa - Angola</option>
											<option>Asia - Vietnam</option>
											<option>Mid East - Dubai</option>
										</select>
									</div>
								</div>
								<div class="holder">
									<label for="date">Select Your Dates</label>
									<div class="select-holder">
										<select class="trip dark" id="date">
											<option>Jan - Mar 2016</option>
											<option>Apr - Jun 2016</option>
											<option>Jul - Sep 2016</option>
											<option>Oct - Dec 2016</option>
										</select>
									</div>
								</div>
								<div class="holder">
									<input class="btn btn-trip" type="submit" value="go wild">
								</div>
							</fieldset>
						</form>
					</div>
				</div>
        <div class="feature-block">
            <div class="holder">
                <ul>

                    <li v-for="category in allCategories" :key="category.id">
                        <a href="#">
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
export default {
    name: "banner",
    computed: mapGetters(['allCategories']),
    methods: mapActions(['getAllCategories']),
    async mounted() {
        this.getAllCategories(8);
    }
}   
</script>

<style scoped>
	.fullscreenabanner {
		height: 820px;
	}
</style>