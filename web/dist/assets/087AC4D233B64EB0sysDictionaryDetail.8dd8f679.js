/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
import{s as e,a,V as l,o as t,c as s,b as i,d as o,w as r,f as u,t as n,F as d,Q as p,dx as c,R as m,j as D,E as y,k as v}from"./087AC4D233B64EB0index.5fcd1422.js";import{E as B}from"./087AC4D233B64EB0dialog.e17c90b5.js";import"./087AC4D233B64EB0overlay.1520877c.js";import{a as f,E as b}from"./087AC4D233B64EB0form-item.69b26929.js";/* empty css                              */import{E as g}from"./087AC4D233B64EB0input-number.f5afd020.js";import{E as C}from"./087AC4D233B64EB0switch.5175e219.js";import{E}from"./087AC4D233B64EB0pagination.81705c9f.js";/* empty css                            */import"./087AC4D233B64EB0select.591b7648.js";import"./087AC4D233B64EB0scrollbar.b45aa0d9.js";/* empty css                               */import{E as j,a as A}from"./087AC4D233B64EB0table-column.de1f83f6.js";import"./087AC4D233B64EB0checkbox.2864820c.js";/* empty css                                *//* empty css                               */import"./087AC4D233B64EB0index.71b3c6ca.js";import"./087AC4D233B64EB0castArray.ef7e3470.js";import"./087AC4D233B64EB0_baseClone.bd09a709.js";import"./087AC4D233B64EB0_Uint8Array.72188c2c.js";import"./087AC4D233B64EB0_initCloneObject.67d0c1fc.js";import"./087AC4D233B64EB0index.c25afd12.js";import"./087AC4D233B64EB0isEqual.c192b418.js";import"./087AC4D233B64EB0index.d064f019.js";import"./087AC4D233B64EB0strings.84714f21.js";import"./087AC4D233B64EB0debounce.679d5750.js";import"./087AC4D233B64EB0_baseIteratee.bf25bc12.js";import"./087AC4D233B64EB0index.508774e7.js";const h=a=>e({url:"/sysDictionaryDetail/createSysDictionaryDetail",method:"post",data:a}),w={class:"gva-table-box"},_={class:"gva-btn-list justify-between"},x=i("span",{class:"text font-bold"},"字典详细内容",-1),V={class:"gva-pagination"},k={class:"dialog-footer"},I=Object.assign({name:"SysDictionaryDetail"},{__name:"sysDictionaryDetail",props:{sysDictionaryID:{type:Number,default:0}},setup(I){const S=I,z=a({label:null,value:null,status:!0,sort:null}),U=a({label:[{required:!0,message:"请输入展示值",trigger:"blur"}],value:[{required:!0,message:"请输入字典值",trigger:"blur"}],sort:[{required:!0,message:"排序标记",trigger:"blur"}]}),q=a(1),T=a(0),F=a(10),L=a([]),M=e=>{F.value=e,N()},O=e=>{q.value=e,N()},N=async()=>{const a=await(l={page:q.value,pageSize:F.value,sysDictionaryID:S.sysDictionaryID},e({url:"/sysDictionaryDetail/getSysDictionaryDetailList",method:"get",params:l}));var l;0===a.code&&(L.value=a.data.list,T.value=a.data.total,q.value=a.data.page,F.value=a.data.pageSize)};N();const Q=a(""),R=a(!1),G=async a=>{const l=await(t={ID:a.ID},e({url:"/sysDictionaryDetail/findSysDictionaryDetail",method:"get",params:t}));var t;Q.value="update",0===l.code&&(z.value=l.data.reSysDictionaryDetail,R.value=!0)},H=()=>{R.value=!1,z.value={label:null,value:null,status:!0,sort:null,sysDictionaryID:S.sysDictionaryID}},J=async a=>{m.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((async()=>{var l;0===(await(l={ID:a.ID},e({url:"/sysDictionaryDetail/deleteSysDictionaryDetail",method:"delete",data:l}))).code&&(D({type:"success",message:"删除成功"}),1===L.value.length&&q.value>1&&q.value--,N())}))},K=a(null),P=async()=>{K.value.validate((async a=>{if(z.value.sysDictionaryID=S.sysDictionaryID,!a)return;let l;switch(Q.value){case"create":default:l=await h(z.value);break;case"update":l=await(t=z.value,e({url:"/sysDictionaryDetail/updateSysDictionaryDetail",method:"put",data:t}))}var t;0===l.code&&(D({type:"success",message:"创建/更改成功"}),H(),N())}))},W=()=>{Q.value="create",R.value=!0};return l((()=>S.sysDictionaryID),(()=>{N()})),(e,a)=>{const l=y,m=j,D=A,h=E,I=v,S=f,N=C,X=g,Y=b,Z=B;return t(),s("div",null,[i("div",w,[i("div",_,[x,o(l,{type:"primary",icon:"plus",onClick:W},{default:r((()=>[u("新增字典项")])),_:1})]),o(D,{ref:"multipleTable",data:L.value,style:{width:"100%"},"tooltip-effect":"dark","row-key":"ID"},{default:r((()=>[o(m,{type:"selection",width:"55"}),o(m,{align:"left",label:"日期",width:"180"},{default:r((e=>[u(n(d(p)(e.row.CreatedAt)),1)])),_:1}),o(m,{align:"left",label:"展示值",prop:"label"}),o(m,{align:"left",label:"字典值",prop:"value"}),o(m,{align:"left",label:"扩展值",prop:"extend"}),o(m,{align:"left",label:"启用状态",prop:"status",width:"120"},{default:r((e=>[u(n(d(c)(e.row.status)),1)])),_:1}),o(m,{align:"left",label:"排序标记",prop:"sort",width:"120"}),o(m,{align:"left",label:"操作",width:"180"},{default:r((e=>[o(l,{type:"primary",link:"",icon:"edit",onClick:a=>G(e.row)},{default:r((()=>[u("变更")])),_:2},1032,["onClick"]),o(l,{type:"primary",link:"",icon:"delete",onClick:a=>J(e.row)},{default:r((()=>[u("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),i("div",V,[o(h,{"current-page":q.value,"page-size":F.value,"page-sizes":[10,30,50,100],total:T.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:O,onSizeChange:M},null,8,["current-page","page-size","total"])])]),o(Z,{modelValue:R.value,"onUpdate:modelValue":a[5]||(a[5]=e=>R.value=e),"before-close":H,title:"create"===Q.value?"添加字典项":"修改字典项"},{footer:r((()=>[i("div",k,[o(l,{onClick:H},{default:r((()=>[u("取 消")])),_:1}),o(l,{type:"primary",onClick:P},{default:r((()=>[u("确 定")])),_:1})])])),default:r((()=>[o(Y,{ref_key:"dialogForm",ref:K,model:z.value,rules:U.value,"label-width":"110px"},{default:r((()=>[o(S,{label:"展示值",prop:"label"},{default:r((()=>[o(I,{modelValue:z.value.label,"onUpdate:modelValue":a[0]||(a[0]=e=>z.value.label=e),placeholder:"请输入展示值",clearable:"",style:{width:"100%"}},null,8,["modelValue"])])),_:1}),o(S,{label:"字典值",prop:"value"},{default:r((()=>[o(I,{modelValue:z.value.value,"onUpdate:modelValue":a[1]||(a[1]=e=>z.value.value=e),placeholder:"请输入字典值",clearable:"",style:{width:"100%"}},null,8,["modelValue"])])),_:1}),o(S,{label:"扩展值",prop:"extend"},{default:r((()=>[o(I,{modelValue:z.value.extend,"onUpdate:modelValue":a[2]||(a[2]=e=>z.value.extend=e),placeholder:"请输入扩展值",clearable:"",style:{width:"100%"}},null,8,["modelValue"])])),_:1}),o(S,{label:"启用状态",prop:"status",required:""},{default:r((()=>[o(N,{modelValue:z.value.status,"onUpdate:modelValue":a[3]||(a[3]=e=>z.value.status=e),"active-text":"开启","inactive-text":"停用"},null,8,["modelValue"])])),_:1}),o(S,{label:"排序标记",prop:"sort"},{default:r((()=>[o(X,{modelValue:z.value.sort,"onUpdate:modelValue":a[4]||(a[4]=e=>z.value.sort=e),modelModifiers:{number:!0},placeholder:"排序标记"},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue","title"])])}}});export{I as default};
