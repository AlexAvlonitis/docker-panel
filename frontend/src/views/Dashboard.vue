<template>
  <CRow v-if=isLoading alignHorizontal=center>
    <CSpinner />
  </CRow>
  <CCard v-else>
    <CCardHeader>
      <h3>Containers</h3>
    </CCardHeader>
    <CCardBody>
      <CRow>
        <CCol lg="3" sm="6">
          <CCallout color="info">
            <small class="text-muted">Count</small><br>
            <strong class="h4">{{ infoData.containers.count }}</strong>
          </CCallout>
        </CCol>
        <CCol lg="3" sm="6">
          <CCallout color="success">
            <small class="text-muted">Running</small><br>
            <strong class="h4">{{ infoData.containers.running }}</strong>
          </CCallout>
        </CCol>
        <CCol lg="3" sm="6">
          <CCallout color="secondary">
            <small class="text-muted">Paused</small><br>
            <strong class="h4">{{ infoData.containers.paused }}</strong>
          </CCallout>
        </CCol>
        <CCol lg="3" sm="6">
          <CCallout color="danger">
            <small class="text-muted">stopped</small><br>
            <strong class="h4">{{ infoData.containers.stopped }}</strong>
          </CCallout>
        </CCol>
      </CRow>
    </CCardBody>
  </CCard>
</template>

<script>
  import InfoData from '@/objects/infoData';
  import HttpClient from '@/utils/httpClient';

  export default {
    name: 'Dashboard',
    data() {
      return {
        infoData: null,
        isLoading: true
      }
    },
    methods: {
      normalizeData: function(infoJSON) {
        return new InfoData(infoJSON)
      }
    },
    created () {
      HttpClient
        .get('info')
        .then((data) => {
          this.infoData = this.normalizeData(data)
          this.isLoading = false;
        })
    }
  }
</script>
