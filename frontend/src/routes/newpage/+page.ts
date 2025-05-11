import { get } from 'svelte/store';
import type { PageLoad } from './$types';


export const load: PageLoad = async () => {

    const response = await fetch('http://localhost:8080/newPageDataHandler')

    const data = await response.json()

    console.log(data)

    return {
        get: {
            message: data.message,
        }
    }
};