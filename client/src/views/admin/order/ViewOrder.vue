<template>
    <div id="page-wrapper" >
        <div id="page-inner">
                <!-- /. ROW  -->
            <div class="row">
                <div class="col-lg-12">
                    <div class="panel panel-default">
                        <div class="panel-body">
                            <div class="row">
                                <div class="col-lg-6">
   
                                    <div class="form-group">
                                        <label>Ід</label>
                                        <input :value="oneOrder.id" readonly class="form-control">
                                    </div>

                                    <div class="form-group">
                                        <label>Опис</label>
                                        <textarea :value="oneOrder.text" readonly class="form-control" rows="3"></textarea>
                                    </div>

                                    <div class="form-group">
                                        <label>Номер телефону</label>
                                        <input :value="oneOrder.phone" readonly class="form-control">
                                    </div>

                                    <div class="form-group">
                                        <label>Клієнт</label>
                                        <input :value="oneOrder.user_name" readonly class="form-control">
                                    </div>

                                    <div class="form-group">
                                        <label>Загальна сума</label>
                                        <input :value="oneOrder.general_price" readonly class="form-control">
                                    </div>

                                    <div class="form-group">
                                        <label>Оплата</label>
                                        <input :value="oneOrder.paid" readonly class="form-control">
                                    </div>

                                    <div class="form-group">
                                        <label>Дата оформлення</label>
                                        <input :value="format_date(oneOrder.created_at, 'DD MMM YYYY hh:ss')" readonly class="form-control">
                                    </div>

                                    <div class="form-group">
                                        <label>Поїздки</label>
                                        <div v-for="trip in oneOrder.trips_orders" :key="trip.id">

                                            <label>Тур: </label>
                                            <router-link :to="{name: 'adminUpdateTour', params: {slug: trip.slug}}"> {{trip.title}}</router-link><p></p>
                                            <label>Дата: </label>
                                            <span> {{format_date(trip.start, 'DD MMM')}} - {{format_date(trip.end, 'DD MMM YYYY')}}</span><p></p>
                                            <label>Кількість: </label>
                                            <span> {{trip.count}}</span><p></p>
                                            <label>Ціна: </label>
                                            <span> {{trip.price}}</span><p></p>
                                            <label>Загальна ціна: </label>
                                            <span> {{trip.price * trip.count}}</span><p></p>
                                            <hr>
                                        </div>
                                    </div>
                         
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
    import moment from 'moment'

    export default {
        name: "ViewOrder",
        
        computed: mapGetters(['oneOrder']),
        props: ['id'],
        methods: {
            ...mapActions(['getOneOrder']),
            format_date(value, format='DD MMM YYYY'){
                if (value) {
                    return moment(String(value)).format(format); 
                }
            },
        },
        async mounted() {
            this.getOneOrder(this.id);
        },
        
    }
</script>


<style scoped>

</style>