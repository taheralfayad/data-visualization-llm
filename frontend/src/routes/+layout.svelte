<script>
    let { children } = $props();
</script>

<nav>
    <div class="link-container">
        <a href="/">
            Home Screen
        </a>
        <a href="/newpage">
            New Page
        </a>
    </div>

</nav>

{@render children()}

<h1>thank you for your attention to this matter</h1>


<style>
    nav {
        background-color: rgb(89, 87, 87);
    }
    a {
        display: flex;
        align-items: center;
        justify-content: center;
        text-decoration: none;
        color: white;
        font-weight: bold;
        padding: 15px 32px;
        border-radius: 4px;
        transition: background-color 0.3s;
        font-family: Arial, Helvetica, sans-serif;
    }
    a:hover {
        background-color: #bb4f15;
        text-decoration: none;
    }
    .link-container {
        display: flex;
        gap: 1rem;
        padding: 1rem;
    }
</style>