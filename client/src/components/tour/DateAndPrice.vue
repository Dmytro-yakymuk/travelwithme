<template>
    <div class="tab-pane active" id="tab3">
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
                                <strong class="date-text">Ціна</strong>
                            </th>
                            <th>
                                &nbsp;
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="trip in this.trips" :key="trip.id">
                            <td><div class="cell"><div class="middle">{{ format_date(trip.start) }}</div></div></td>
                            <td><div class="cell"><div class="middle">{{ format_date(trip.end) }}</div></div></td>
                            <td><div class="cell"><div class="middle">{{ trip.price }}</div></div></td>
                            <td><div class="cell"><div class="middle">
                                <a href="#" @click="goOrder(trip.id)" class="btn btn-default">Замовити</a>
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
</template>

<script>
    import moment from 'moment'   
    import {mapGetters} from "vuex"

    export default {
        name: 'dateAndPrice',
        props: ['trips'],
        computed: mapGetters(['loggedIn']),
        methods: { 
            format_date(value){
                if (value) {
                    return moment(String(value)).format('DD MMM YY'); 
                }
            },
            goOrder(id) {
                if (this.loggedIn && localStorage.getItem('role') == 'client') {
                    this.$store.dispatch('goOrder', id)
                } else {
                    this.$router.push({ name: 'login'})
                }
                
            }
        },
    }
</script>