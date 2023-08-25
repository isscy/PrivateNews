import {defineStore} from 'pinia'

                                            // id      对象
export const useCounterStore = defineStore("counter", {
    state() {
        return {
            count: 0,
        }
    },
    getters: {
        doubleCount(state){
            return state.count*2;
        }
    },
    actions: {
        increment(){
            this.count ++ ;
        },
    }
})