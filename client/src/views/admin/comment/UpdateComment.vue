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
                                            <label>Повідомлення</label>
                                            <textarea v-model="oneComment.message" class="form-control" rows="5"></textarea>
                                        </div>

                                        <div class="form-group">
                                            <label>Зірки</label>
                                            <input type="number" v-model.number="oneComment.star" min="1" max="5">
                                        </div>

                                        <div class="form-group">
                                            <label>Тур</label>
                                            <select v-model="oneComment.tour_id" class="form-control">
                                                <option 
                                                    v-for="tour in allTours" :key="tour.id"
                                                    :value="tour.id"
                                                    :selected="tour.id == oneComment.tour_id"
                                                >{{tour.title}}</option>
                                            </select>
                                        </div>
  
                                        <button type="submit" class="btn btn-default">Редагувати</button>
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
        name: "AdminUpdateComment",
        props: ['id'],
        computed: mapGetters(['oneComment', 'getSuccessMessage', 'getFailMessage', 'allTours']),
        methods: {
            ...mapActions(['getOneComment', 'updateComment', 'getAllTours']),

            submitForm() {
                this.updateComment({
                    'id': this.oneComment.id,
                    'message': this.oneComment.message,
                    'star': this.oneComment.star,
                    'tour_id': this.oneComment.tour_id
                });
            },
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