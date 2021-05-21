<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            <div class="row">
                <div class="col-md-12">

                    <div class="panel panel-default">
                        <div class="panel-heading"></div>
                        <div class="panel-body">
                            <div class="table-responsive">
                                <table class="table table-striped table-bordered table-hover" id="dataTables-example">
                                    <thead>
                                        <tr>
                                            <th>Ід</th>
                                            <th>Коментарій</th>
                                            <th class="text-center">Телефон</th>
                                            <th class="text-center">Клієнт</th>
                                            <th class="text-center">Загальна сума</th>
                                            <th class="text-center">Оплачено</th>
                                            <th class="text-center">Дата оформлення</th>
                                            <th class="text-center"></th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        
                                        <tr v-for="order, k in allOrders" :key="order.id" class="gradeX" >
                                            
                                            <td>
                                                {{ k+1 }}
                                            </td>
                                            <td>
                                                {{ order.text }}
                                            </td>
                                            <td class="text-center">
                                                {{ order.phone }}
                                            </td>
                                            <td class="text-center">
                                                {{ order.user_name }}
                                            </td>
                                            <td class="text-center">
                                                {{ order.general_price }}
                                            </td>
                                            <td class="text-center">
                                                <!-- <input type="checkbox" v-model="order.paid" @click="paid(order.id, order.paid)"> -->
                                                <input type="checkbox" v-model="order.paid" onclick="return false;">
                                            </td>
                                            <td class="text-center">
                                                {{ format_date(order.created_at, 'DD MMM YYYY hh:ss') }}
                                            </td>
                                            <td class="text-center">
                                                <a class="btn btn-primary btn-action" @click="goTo(order.id)"><i class="fa fa-book"></i></a>
                                                <a v-if="order.paid == true" class="btn btn-primary btn-action" @click="goDownload(order)"><i class="fa fa-download" aria-hidden="true"></i></a>
                                                <a v-if="order.paid == false && role == 'client'" class="btn btn-primary btn-action" @click="goPaid(order)"><i class="fa fa-credit-card" aria-hidden="true"></i></a>
                                            </td>
                                        </tr>

                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import {mapGetters, mapActions, mapMutations} from "vuex"
    import moment from 'moment'

    export default {
        name: "AdminOrders",
        data: function () {
            return {
                role: localStorage.getItem('role'),
            }
        },
        computed: mapGetters(['allOrders']),
        methods: {
            ...mapActions(['getAllOrders', 'updatePaid', 'createOrder', 'getAllTripsWhichAttach', 'downloadOrder']),
            ...mapMutations(['updateAttachTrips']),
            goTo(id) {
                this.$router.push({ name: 'adminViewOrder', params: {id: id} })
            },
            goDownload(order) {
                this.downloadOrder({
                    'id': order.id,
                }).then (response => {
                   if (response.data.result) {
                        const blob = new Blob([response.data.result], { type: 'application/pdf' })
                        const link = document.createElement('a')
                        link.href = URL.createObjectURL(blob)
                        link.download = label
                        link.click()
                        URL.revokeObjectURL(link.href)
                    } else {
                        alert("a;a;a;")
                    }
                })
            },
            // goDelete(slug) {
            //     this.$router.push({ name: 'adminDeleteTour', params: {slug: slug} })
            // },
            // paid(id, paid) {
            //     this.updatePaid({
            //         'id': id,
            //         'paid': paid,
            //     })
            // },
            goPaid(order) {
                this.createOrder({
                    'id': order.id,
                    'phone': order.phone,
                    'text': order.text,
                    'totalPrice': order.general_price
                }).then (response => {
                    
                    (async function() {
                        return window.open(response.data.result)
                    })(); 
                    this.updateAttachTrips(this.getAllTripsWhichAttach());
                    // this.$router.push({ name: 'cart'})
                })
            },
            format_date(value, format='DD MMM YYYY'){
                if (value) {
                    return moment(String(value)).format(format); 
                }
            },
        },
        async mounted() {
            this.getAllOrders();
        },
        
        
    }
</script>

<style scoped>

</style>