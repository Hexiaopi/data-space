import router from '@/router'
import { Directive, DirectiveBinding, unref } from 'vue'

const auth: Directive = {
    mounted(el: HTMLElement, binding: DirectiveBinding) {
        const currentRoute = unref(router.currentRoute)
        const btns = currentRoute.meta?.btns?.map((item: any) => item.name) || []
        if (!btns.includes(binding.value)) {
            el.remove()
        }
    }
}

export default auth