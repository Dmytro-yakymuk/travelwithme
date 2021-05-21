<template>
    <div class="tab-pane active" id="tab3">

        <div v-if="getSuccessMessage != ''" class="alert alert-success" role="alert">{{getSuccessMessage}}</div>
        <div v-if="getFailMessage != ''" class="alert alert-danger" role="alert">{{getFailMessage}}</div>

        <div class="table-container">
            <div class="table-responsive">
                <table class="table table-striped" v-if="oneTour.trips && oneTour.trips.length > 0">
                    <thead>
                        <tr>
                            <th>
                                <strong class="date-text">Початок</strong>
                            </th>
                            <th>
                                <strong class="date-text">Завершення</strong>
                            </th>
                            <th>
                                <strong class="date-text">Кількість місць</strong>
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
                        <tr v-for="trip in oneTour.trips" :key="trip.id">
                            <td><div class="cell"><div class="middle">{{ format_date(trip.start) }}</div></div></td>
                            <td><div class="cell"><div class="middle">{{ format_date(trip.end) }}</div></div></td>
                            <td><div class="cell"><div class="middle">{{ trip.freeCount }} із {{ trip.count }}</div></div></td>
                            <td><div class="cell"><div class="middle">{{ trip.price }} грн</div></div></td>
                            <td><div class="cell"><div class="middle">
                                <a v-if="trip.freeCount >= 1" @click="attach(trip.id)" class="btn btn-default">Замовити</a>
                            </div></div></td>
                        </tr>
                    </tbody>
                </table>
                <div v-else>
                    <p>Немає поїздок</p>
                </div>
            </div>
            <!-- <div class="load-more text-center text-uppercase">
                <a href="#">Більше</a>
            </div> -->
        </div>
    </div>
</template>

<script>
    import moment from 'moment'   
    import {mapGetters} from "vuex"

    export default {
        name: 'dateAndPrice',
        computed: mapGetters(['loggedIn', 'oneTour', 'getSuccessMessage', 'getFailMessage']),
        methods: { 
            format_date(value){
                if (value) {
                    return moment(String(value)).format('DD MMM YY'); 
                }
            },
            attach(id) {
                if (this.loggedIn && localStorage.getItem('role') == 'client') {
                    this.$store.dispatch('attach', id)
                        .then(response => {
                            //console.log(response)
                        } 
                            // this.$router.push({ name: 'cart'})
                        )
                        .catch(error => {
                            if (error.response != null) {
                                if (error.response.status === 401) {
                                    this.$router.push({ name: 'logout' })
                                }
                            }
                        })  
                } else if(this.loggedIn && localStorage.getItem('role') != 'client') {
                    alert('Only client can to order trip')
                } else {
                    this.$router.push({ name: 'login'})
                }
                
            }
        },
    }
</script>