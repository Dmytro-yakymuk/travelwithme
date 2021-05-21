<template>
    <div class="tab-pane active" id="tab2">
        <div class="row">
            
            <div class="col-md-12">
                <div class="header-box">
                    <!-- <a href="#" class="link-right">Залишити відгук</a> -->
                    <span class="rate-left">
                        <strong class="title">Загальний рейтинг</strong>
                        <span class="star-rating">
                            <span v-for="star in 5" :key="star" :class="oneTour.generalStar >= star ? '' : 'disable'"><span class="icon-star"></span></span>
                        </span>
                        <span class="value">{{oneTour.generalStar}}/5</span>
                    </span>
                </div>
                <div class="comments reviews-body">
                    <div v-if="oneTour.comments.length > 0 " class="comment-holder">

                        <div v-for="commnet in oneTour.comments" :key="commnet.id" class="comment-slot">
                            <div class="thumb">
                                <a href="#"><img src="../../assets/img/thumbs/avatar_male.png" height="50" width="50" alt="image description"></a>
                            </div>
                            <div class="text">
                                <header class="comment-head">
                                    <div class="left">
                                        <strong class="name">
                                            <a href="#">{{ commnet.user_name }}</a>
                                        </strong>
                                        <div class="meta">Опубліковано {{ format_date(commnet.created_at) }}</div>
                                    </div>
                                    <div class="right">
                                        <div class="star-rating">
                                            <span v-for="star in 5" :key="star" :class="commnet.star >= star ? '' : 'disable'"><span class="icon-star"></span></span>
                                        </div>
                                        <span class="value">{{ commnet.star }}/5</span>
                                    </div>
                                </header>
                                <div class="des">
                                    <p>{{ commnet.message }}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <!-- <div class="link-more text-center">
                        <a href="#">Переглянути більше - х відгуків</a>
                    </div> -->
                </div>
            </div>
        </div>
    </div>

</template>

<script>
export default {
    name: 'comments'
}
</script>

<script>
    import moment from 'moment'   
    import {mapGetters} from "vuex"

    export default {
        name: 'comments',
        computed: mapGetters(['loggedIn', 'oneTour', 'getTotalCountStar']),
        methods: { 
            format_date(value){
                if (value) {
                    return moment(String(value)).format('DD MMM YY'); 
                }
            },
            
        },
    }
</script>