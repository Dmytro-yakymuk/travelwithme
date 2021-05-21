<template>
    <div id="page-wrapper" >
        <div id="page-inner">
            <div class="row">
                <div class="col-lg-12">
                    <div class="panel panel-default">
                        <div class="panel-body">
                            <div class="row">
                                <div class="col-lg-6">
                                    <div class="form-group">
                                        <label>Повідомлення</label>
                                        <textarea v-model="oneComment.message" class="form-control" rows="5" readonly></textarea>
                                    </div>
                                    <div class="form-group">
                                        <label>Зірки</label>
                                        <input type="number" v-model.number="oneComment.star" min="1" max="5" readonly>
                                    </div>
                                    <div class="form-group">
                                        <label>Тур</label>
                                        <select v-model="oneComment.tour_id" class="form-control" disabled>
                                            <option 
                                                v-for="tour in allTours" :key="tour.id"
                                                :value="tour.id"
                                                :selected="tour.id == oneComment.tour_id"
                                            >{{tour.title}}</option>
                                        </select>
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

    export default {
        name: "AdminViewComment",
        props: ['id'],
        computed: mapGetters(['oneComment', 'allTours']),
        methods: {
            ...mapActions(['getOneComment', 'getAllTours']),
        },
        async mounted() {
            this.getOneComment(this.id);
            this.getAllTours({
                'order_by': "id desc"
            });
        }
    }
</script>

<style scoped></style>