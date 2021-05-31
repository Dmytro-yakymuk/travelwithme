import _ from 'lodash'

export default {
    data() {
        return {
            page: +this.$route.query.page || 1,
            pageSize: 2,
            allItems: [],
        }
    },
    methods: {
        pageChangeHandler(page) {
            this.$router.push({ path: this.$route.path, query: {
                category_id: this.$route.query.category_id,
                region_id: this.$route.query.region_id,
                price: this.$route.query.price,
                page: page, 
            } })
            this.updateToursMutation(this.allItems[page - 1] || this.allItems[0])
            this.updateCountTours(this.allTours.length)
        },
        setupPagination(allItems) {
            // console.log(allItems)
            // console.log('getPageCount1 = '+this.getPageCount)
            var size = _.size(allItems)
            
            if (size > 0) {
                this.allItems = _.chunk(allItems, this.pageSize)
                this.updateToursMutation(this.allItems[this.page - 1] || this.allItems[0])
            } else {
                this.updateToursMutation([])
            }

            this.updatePageCount(size)
            this.updateCountTours(size)

            // console.log(allItems)
            // console.log('getPageCount2 = '+this.getPageCount)
            // console.log('//////////////////////////////')
        }
    }
}