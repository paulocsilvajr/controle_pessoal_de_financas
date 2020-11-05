<template>
    <p class="dev">
        Desenvolvido por <a href="https://github.com/paulocsilvajr">
        Paulo C. Silva Jr
        <i class="fab fa-github" aria-hidden="true"></i></a>
    </p>
</template>

<script>
    export default {
        mounted() {
            console.log('Componente desenvolvedor mountado')
        }
    }
</script>

<style scoped>
    .dev {
        margin-top: 2em;
        text-align: center;
    }
</style>
