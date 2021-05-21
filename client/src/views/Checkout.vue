<template>
    <main id="main">
        <!-- top information area -->
        <div class="inner-top">
            <div class="container">
                <h1 class="inner-main-heading">Оформлення замовлення</h1>
                <!-- breadcrumb -->
                <nav class="breadcrumbs">
                    <ul>
                        <li><a href="#">Home</a></li>
                        <li><a href="#">pages</a></li>
                        <li><span>Checkout</span></li>
                    </ul>
                </nav>
            </div>
        </div>
        <div class="inner-main common-spacing container">
            
            <form @submit.prevent="submitForm" class="booking-form" action="#">
                <div class="row">
                    <div class="col-md-10">
                        <div class="form-holder">
                            <h2 class="small-size">Додаткові нотатки</h2>
                             <div class="row">
                                <div class="col-md-4">
                                    <div class="hold">
                                        <label for="phn">Номер телефону*</label>
                                        <input v-model="oneUser.phone" type="text" id="phn" class="form-control">
                                    </div>
                                </div>
                            </div>
                            <div class="wrap">
                                <div class="hold">
                                    <label for="txt">Ваш коментарій</label>
                                    <textarea v-model="text" id="txt" class="form-control" placeholder="Будь ласка, введіть сюди будь-яку додаткову інформацію. "></textarea>
                                </div>
                            </div>
                            <div class="order-block">
                                <h2 class="small-size">Попередній перегляд замовлення</h2>
                                <div class="wrap">
                                    <table class="product-table">
                                        <thead>
                                            <tr>
                                                <th>Вибрані тури</th>
                                                <th>Загальна сума</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            <tr v-for="trip in allTripsWhichAttach" :key="trip.id">
                                                <td>
                                                    <span class="title"><a :href="'/tour/'+trip.slug">{{ trip.title }}</a> &emsp;<span class="amount">x&emsp; {{ trip.count }}</span></span>
                                                    <time>{{format_date(trip.start, 'DD MMM')}} - {{format_date(trip.end, 'DD MMM YYYY')}}</time>
                                                </td>
                                                <td>
                                                    <span class="amount">{{ trip.price * trip.count }} грн</span>
                                                </td>
                                            </tr>

                                        </tbody>
                                        <tfoot>
                                            <tr>
                                                <td>Загальна сума</td>
                                                <td>{{getTotalPrice}} грн</td>
                                            </tr>
                                        </tfoot>
                                    </table>
                                    <button type="submit" class="btn btn-default">Оплатити</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    </main>
</template>


<script>
    import {mapGetters, mapActions, mapMutations} from "vuex"
    import moment from 'moment' 

    export default { 
        data: function () {
            return {
                text: ''
            }
        },
        name: 'Checkout',
        computed: mapGetters(['loggedIn', 'allTripsWhichAttach', 'getTotalPrice', 'oneUser']),
        methods: { 
            ...mapActions(['getAllTripsWhichAttach', 'deleteTrip', 'createOrder', 'getOneUser']),
            ...mapMutations(['updateAttachTrips']),
            format_date(value, format='DD MMM YYYY'){
                if (value) {
                    return moment(String(value)).format(format); 
                }
            },
            submitForm() {
                this.createOrder({
                    'id': "",
                    'phone': this.oneUser.phone,
                    'text': this.text,
                    'totalPrice': this.getTotalPrice
                }).then (response => {
                    (async function() {
                        return window.open(response.data.result)
                    })(); 
                    this.updateAttachTrips(this.getAllTripsWhichAttach());
                    this.$router.push({ name: 'cart'})
                })
            }
        },
        async mounted() {
            this.getOneUser();
        }
    }
</script>
