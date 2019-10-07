<script>
  export let fileuri;
  
  let filepath = "/assets/sample-image.jpeg"
  let filename = "sample-image.jpg"
  let filesize = 65000
  let fileLastMod = "1 Jan, 2017"

  const fmtByte = (byteSize) => {
    if (byteSize / 1000000 > 1) {
      return `${byteSize/1000000}MB`
    } else if (byteSize / 1000 > 1) {
      return `${byteSize/1000}KB`
    } else {
      return `${byteSize}B`
    }
  }

</script>

<style>
  .image-wrapper{
    width: 100%;
    height: 500px;
    border: 1px solid #ddd;
  }
  .image-wrapper img {
    object-fit: cover;
    min-width: 100%;
    min-height: 100%;
    width: auto;
    height: auto;
    max-width: 100%;
    max-height: 100%;
  }
</style>

<div class="card">
  <div class="card-image">
    <div class="image-wrapper">
      <img src={fileuri} alt="" class="img-responsive">
    </div>
  </div>
</div>

<div class="card" style="margin-top: 5px; padding: 5px 0px 5px 20px;">
  <div class="tile tile-centered">
    <div class="tile-content">
      <div class="tile-title">{filename}</div>
      <small class="tile-subtitle text-gray">{fmtByte(filesize)} · {fileLastMod} · {filepath}</small>
    </div>
    <div class="tile-action">
      <button class="btn btn-link">
        <i class="icon icon-more-vert"></i>
      </button>
    </div>
  </div>
</div>