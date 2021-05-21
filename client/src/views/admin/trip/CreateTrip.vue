<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            
            <div v-if="getSuccessMessage != ''" class="alert alert-success" role="alert">{{getSuccessMessage}}</div>
            <div v-if="getFailMessage != ''" class="alert alert-danger" role="alert">{{getFailMessage}}</div>

            <div class="row">
                <div class="col-lg-12">
                    <div class="panel panel-default">
                        <div class="panel-body">
                            <div class="row">
                                <div class="col-lg-6">
                                    <form role="form" @submit.prevent="submitForm">
                                        
                                        <div class="form-group">
                                            <label>Початок</label>
                                            <input type="date" v-model="start" class="form-control"
                                                value=""
                                                min="2021-01-01" max="2022-01-01">
                                        </div>

                                        <div class="form-group">
                                            <label>Кінець</label>
                                            <input type="date" v-model="end" class="form-control"
                                                value=""
                                                min="2021-01-01" max="2022-01-01">
                                        </div>

                                        <div class="form-group">
                                            <label>Ціна</label>
                                            <input type="number" v-model.number="price" min="0">
                                        </div>

                                        <div class="form-group">
                                            <label>Кількість</label>
                                            <input type="number" v-model.number="count" min="0">
                                        </div>
 

                                        <div class="form-group">
                                            <label>Тур</label>
                                            <select v-model="tour_id"  class="form-control">
                                                <option 
                                                    v-for="tour in allTours" :key="tour.id"
                                                    :value="tour.id"
                                                    :selected="tour.id == tour_id"
                                                >{{tour.title}}</option>
                                            </select>
                                        </div>
                                    
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

    export default {
        name: "AdminCreateTrip",
        data: function () {
            return {
                start: '',
                end: '',
                price: 0,
                count: 0,
                tour_id: 0,
            }
        },

        computed: mapGetters(['allTours', 'getSuccessMessage', 'getFailMessage']),
        methods: {
            ...mapActions(['getAllTours', 'createTrip']),

            submitForm() {
                this.createTrip({
                    'start': this.format_date(this.start),
                    'end': this.format_date(this.end),
                    'price': this.price,
                    'count': this.count,
                    'tour_id': this.tour_id
                })
                .then(response => {
                    //this.$router.push({ name: 'adminGetTours'})
                })
                .catch(error => {
                    if (error.response != null) {
                        if (error.response.status === 401) {
                            this.$router.push({ name: 'logout' })
                        }
                    }
                })  
            },

            format_date(value){
                var datum = Date.parse(value);
                return datum/1000;
			},
        },
        async mounted() {
            this.getAllTours({
                'user_id': localStorage.getItem('id'),
                'order_by': "id desc"
            });
        },
        
    }
</script>

<style scoped>

</style>