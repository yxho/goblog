<template>
  <div>
    <h3>用户列表页面</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6"><a-input-search placeholder="输入用户名查找" enter-button></a-col>
        <a-col :span="4">
          <a-button type="primary">新增</a-button>
        </a-col>
      </a-row>
      <a-table :pagination='paginationOption' :dataSource="userlist"></a-table>
    </a-card>
  </div>
</template>

<script>
const columns = []
export default {
data(){
  return{
    paginationOption:{
      pageSizeOptions:['2','5','10'],
      defaultCurrent:1,
      defaultPageSize:2,
      total:0,
      showSizeChanger:true,
      showTotal:(total)=>`共${total}条`,
      onChange:(page,pageSize)=>{},
      onShowSizeChange:(current,size)=>{},
    },
    userlist:[],
    columns,
  }
},
created(){},
methods:{
  getUserList(){
    this.$http.get('users',{params:{pagesize:this.paginationOption.defaultPageSize,
    pagenum:this.paginationOption.defaultCurrentFF}})
  }
}
}
</script>
