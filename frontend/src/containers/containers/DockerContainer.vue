<template>
  <div v-on:click="emitClickEvent" style="cursor: pointer;">
    <CWidgetDropdown
      :header=container.name
      :text=container.image
      color="info"
    >
      <CDropdown
        class="float-right"
        color="transparent p-0"
        placement="bottom-end"
      >
        <template #toggler-content>
          <CIcon name="cil-settings"/>
        </template>
        <CDropdownItem>Action</CDropdownItem>
        <CDropdownItem>Another action</CDropdownItem>
        <CDropdownItem disabled>Disabled action</CDropdownItem>
      </CDropdown>
    </CWidgetDropdown>
  </div>
</template>

<script>
  export default {
    name: 'DockerContainer',
    props: ['container'],
    methods: {
      emitClickEvent: function() {
        this.$emit('container-clicked', this.container)
      }
    }
  }
</script>
