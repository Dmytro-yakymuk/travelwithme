<template>
    <main id="main">
        
        <!-- top information area -->
        <div class="inner-top">
            <div class="container">
                <h1 class="inner-main-heading">Корзина</h1>
                <!-- breadcrumb -->
                <nav class="breadcrumbs">
                    <ul>
                        <li><a href="/">Головна</a></li>
                        <li><span>Корзина</span></li>
                    </ul>
                </nav>
            </div>
        </div>
        <div class="inner-main common-spacing container">
            
            <div v-if="getSuccessMessage != ''" class="alert alert-success" role="alert">{{getSuccessMessage}}</div>
            <div v-if="getFailMessage != ''" class="alert alert-danger" role="alert">{{getFailMessage}}</div>

            <div v-if="allTripsWhichAttach.length > 0" class="cart-holder table-container">
                <div class="table-responsive">
                    <table class="table table-hover table-align-right">
                        <thead>
                            <tr>
                                <th>
                                    &nbsp;
                                </th>
                                <th>
                                    <strong class="date-text">Вибрані тури</strong>
                                    <span class="sub-text">Затверджені дати</span>
                                </th>
                                <th>
                                    <strong class="date-text">Ціна</strong>
                                    <span class="sub-text">Оновлено сьогодні</span>
                                </th>
                                <th>
                                    <strong class="date-text">Кількість людей</strong>
                                    <span class="sub-text">Включаючи дітей після 6 років</span>
                                </th>
                                <th>
                                    <strong class="date-text">Загальна сума</strong>
                                    <span class="sub-text">Без урахування рейсів</span>
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="trip in allTripsWhichAttach" :key="trip.id">
                                <td>
                                    <div class="cell">
                                        <div class="middle">
                                            <a class="delete" @click="destroy(trip.id)"><span class="icon-trash"></span></a>
                                        </div>
                                    </div>
                                </td>
                                <td>
                                    <div class="cell">
                                        <div class="middle">
                                            <div class="info">
                                                
                                                <div class="text-wrap">
                                                    <strong class="product-title"><a :href="'/tour/'+trip.slug">{{ trip.title }}</a></strong>
                                                    <time class="time">{{format_date(trip.start, 'DD MMM')}} - {{format_date(trip.end, 'DD MMM YYYY')}}</time>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </td>
                                <td>
                                    <div class="cell">
                                        <div class="middle">
                                            <span class="price">{{ trip.price }} грн</span>
                                        </div>
                                    </div>
                                </td>
                                <td>
                                    <div class="cell">
                                        <div class="middle">
                                            <div class="num-hold">
                                                <a @click="decr(trip.id)" class="minus control"><span class="icon-minus-normal"></span></a>
                                                <a @click="incr(trip.id)" class="plus control"><span class="icon-plus-normal"></span></a>
                                                <span class="val">{{ trip.count }}</span>
                                            </div>
                                        </div>
                                    </div>
                                </td>
                                <td>
                                    <div class="cell">
                                        <div class="middle">
                                            <span class="price">{{ trip.price * trip.count }} грн</span>
                                        </div>
                                    </div>
                                </td>
                            </tr>

                        </tbody>
                    </table>
                </div>
                <div class="cart-option">
                    <div class="button-hold">
                        <a href="/checkout" class="btn btn-default">Оформлення замовлення</a>
                    </div>
                </div>
            </div>
            <div v-else class="content-message">
                <div class="alert alert-info" role="alert">
                    <p>В корзині пусто</p>
                </div>
            </div>
        </div>
    </main>
</template>

<script>
import {mapGetters, mapActions} from "vuex"
import moment from 'moment' 

export default { 
    name: 'Cart',
    computed: mapGetters(['loggedIn', 'allTripsWhichAttach', 'getSuccessMessage', 'getFailMessage']),
    methods: { 
        ...mapActions(['getAllTripsWhichAttach', 'deleteAttachTrip']),
        format_date(value, format='DD MMM YYYY'){
            if (value) {
                return moment(String(value)).format(format); 
            }
        },
        destroy(id) {
            this.deleteAttachTrip(id)
        },
        incr(id) {
            if (this.loggedIn && localStorage.getItem('role') == 'client') {
                this.$store.dispatch('incr', id)
            } else {
                this.$router.push({ name: 'login'})
            }
        },
        decr(id) {
            if (this.loggedIn && localStorage.getItem('role') == 'client') {
                this.$store.dispatch('decr', id)
            } else {
                this.$router.push({ name: 'login'})
            }
            
        }
    },
    async mounted() {
        this.getAllTripsWhichAttach();

    }
}
</script>

<style scoped>
    .content-message {
        min-height: 141px;
    }
</style>