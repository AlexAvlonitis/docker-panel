<template>
  <div v-on:click="emitClickEvent" style="cursor: pointer;">
    <CWidgetDropdown
      :header=network.name
      :text=network.ipAddr
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
    name: 'DockerNetwork',
    props: ['network'],
    methods: {
      emitClickEvent: function() {
        this.$emit('network-clicked', this.network)
      }
    }
  }
</script>
