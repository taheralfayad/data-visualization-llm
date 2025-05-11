<script lang="ts">
    import type { PageProps } from './$types';

    let { data }: PageProps = $props();
</script>

<h1>hello, this is the route newpage</h1>

<p>{data.get.message}</p>