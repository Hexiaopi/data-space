import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
    state: () => ({
        collapsed: false,
    }),
    getters: {
        //
    },
    actions: {
        switchCollapsed() {
            this.collapsed = !this.collapsed
        },
        setCollapsed(collapsed: boolean) {
            this.collapsed = collapsed
        },
    },
})

export default useAppStore