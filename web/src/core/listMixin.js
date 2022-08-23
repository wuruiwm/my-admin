import service from "@/core/request";

export default {
    data() {
        return {
            list: {
                url: "",
                page: 1,
                page_size: 10,
                total: 0,
                data: [],
                date: [],
                param: {
                    keyword: "",
                    start_time: "",
                    end_time: ""
                }
            }
        }
    },
    methods: {
        async getList() {
            if (!this.list.url) {
                console.error("listUrl不能为空")
                return
            }
            let listRes = await service({
                url: this.list.url,
                method: "get",
                params: {page: this.list.page, page_size: this.list.page_size, ...this.list.param}
            })
            if (listRes.code === 0) {
                this.list.data = listRes.data.list
                this.list.total = listRes.data.total
            }
        },
        pageSizeChange(val) {
            this.list.page_size = val
            this.getList()
        },
        pageChange(val) {
            this.list.page = val
            this.getList()
        },
    },
    watch: {
        "list.date": {
            handler(v) {
                if(v){
                    if(v.length === 2){
                        this.list.param.start_time = v[0]
                        this.list.param.end_time = v[1]
                    }
                }else{
                    this.list.param.start_time = ""
                    this.list.param.end_time = ""
                }
            }
        }
    }
}
