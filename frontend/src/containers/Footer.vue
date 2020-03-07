<template>
  <CFooter :fixed="false">
    <div>
      <a href="https://github.com/alexavlonitis/docker-panel" target="_blank">Docker Admin Panel</a>
      <span class="ml-1">&copy; {{new Date().getFullYear()}} Alex Avlonitis</span>
    </div>
    <div class="ml-auto">
      <span class="mr-1">Template by</span>
      <a href="https://coreui.io/vue" target="_blank">CoreUI for Vue</a>
    </div>
  </CFooter>
</template>

<script>
export default {
  name: 'Footer'
}
</script>
