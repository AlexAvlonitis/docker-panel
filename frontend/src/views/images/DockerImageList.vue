<template>
  <CRow>
    <CCol lg="5" md="12" sm="12" overflow-auto>
      <h1>Images List</h1>
      <div v-if=isLoading >
        <CRow alignHorizontal=center>
          <CSpinner />
        </CRow>
      </div>
      <div v-else>
        <CRow>
          <CCol lg="12" class="mb-1">
            Filter by name: <input v-model="imgName"/>
          </CCol>
        </CRow>
        <div v-for="img in imageList()" v-bind:key="img.id">
          <DockerObject :object=img @object-clicked="renderImage"/>
        </div>
      </div>
    </CCol>
    <CCol lg="7" sm="12" md="12">
      <div v-if=selectedImage>
        <h3 class="mb-3">Details</h3>
        <DockerObjectInspect
          httpEndpoint="images"
          :object="selectedImage"
          :isLoading="true"
        />
      </div>
    </CCol>
  </CRow>
</template>

<script>
  import Image from '@/objects/image';
  import HttpClient from '@/utils/httpClient';
  import { arrayFilter } from '@/utils/textUtils'
  import DockerObjectInspect from '@/containers/DockerObjectInspect.vue';
  import DockerObject from '@/containers/DockerObject.vue';

  export default {
    name: 'DockerImageList',
    components: {
      DockerObject,
      DockerObjectInspect
    },
    data() {
      return {
        dockerImages: null,
        isLoading: true,
        selectedImage: null,
        imgName: null
      }
    },
    methods: {
      imageList: function() {
        if (this.imgName == null)
          return this.dockerImages

        return arrayFilter(this.dockerImages, this.imgName)
      },
      normalizeData: function(imagesJSON) {
        return imagesJSON.map(obj => new Image(obj))
      },
      renderImage: function(img) {
        this.selectedImage = img
      }
    },
    created () {
      HttpClient
        .get('images')
        .then((data) => {
          this.dockerImages = this.normalizeData(data)
          this.isLoading = false;
        })
    }
  }
</script>
