/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
import{a,o as e,c as t,d as l,b as o,w as r,f as i,F as u,df as d,g as s,e as n,R as m,j as p,E as y,I as h,k as v}from"./087AC4D233B64EB0index.5fcd1422.js";import{E as c}from"./087AC4D233B64EB0drawer.6a894563.js";import"./087AC4D233B64EB0overlay.1520877c.js";import{E as B,a as f}from"./087AC4D233B64EB0tab-pane.20a0e58d.js";import{E as C}from"./087AC4D233B64EB0dialog.e17c90b5.js";import{a as I,E}from"./087AC4D233B64EB0form-item.69b26929.js";/* empty css                              */import{E as A}from"./087AC4D233B64EB0radio.95689ca2.js";/* empty css                               *//* empty css                            */import"./087AC4D233B64EB0checkbox.2864820c.js";import"./087AC4D233B64EB0scrollbar.b45aa0d9.js";import{E as b,a as D}from"./087AC4D233B64EB0table-column.de1f83f6.js";/* empty css                                *//* empty css                               */import{g,d as j,c as w,u as _,a as k}from"./087AC4D233B64EB0authority.2bc42e73.js";import N from"./087AC4D233B64EB0menus.d55fc980.js";import V from"./087AC4D233B64EB0apis.9cdebe1a.js";import x from"./087AC4D233B64EB0datas.cf632aad.js";import{_ as U}from"./087AC4D233B64EB0warningBar.a4d9a0b0.js";import"./087AC4D233B64EB0index.71b3c6ca.js";import"./087AC4D233B64EB0strings.84714f21.js";import"./087AC4D233B64EB0index.0de793c7.js";import"./087AC4D233B64EB0castArray.ef7e3470.js";import"./087AC4D233B64EB0_baseClone.bd09a709.js";import"./087AC4D233B64EB0_Uint8Array.72188c2c.js";import"./087AC4D233B64EB0_initCloneObject.67d0c1fc.js";import"./087AC4D233B64EB0isEqual.c192b418.js";import"./087AC4D233B64EB0arrays.2c626d3b.js";import"./087AC4D233B64EB0cloneDeep.1ed46a18.js";import"./087AC4D233B64EB0index.d064f019.js";import"./087AC4D233B64EB0index.508774e7.js";import"./087AC4D233B64EB0debounce.679d5750.js";import"./087AC4D233B64EB0_baseIteratee.bf25bc12.js";import"./087AC4D233B64EB0tree.0c7e437e.js";import"./087AC4D233B64EB0index.fc2f110f.js";import"./087AC4D233B64EB0authorityBtn.0274dde2.js";import"./087AC4D233B64EB0api.a9a90911.js";const q={class:"authority"},z={class:"gva-table-box"},F={class:"gva-btn-list"},R={class:"dialog-footer"},S=Object.assign({name:"Authority"},{__name:"authority",setup(S){const O=a([{authorityId:0,authorityName:"根角色"}]),T=a(!1),G=a("add"),K=a({}),P=a("新增角色"),Q=a(!1),$=a(!1),H=a({}),J=a({authorityId:0,authorityName:"",parentId:0}),L=a({authorityId:[{required:!0,message:"请输入角色ID",trigger:"blur"},{validator:(a,e,t)=>/^[0-9]*[1-9][0-9]*$/.test(e)?t():t(new Error("请输入正整数")),trigger:"blur",message:"必须为正整数"}],authorityName:[{required:!0,message:"请输入角色名",trigger:"blur"}],parentId:[{required:!0,message:"请选择父角色",trigger:"blur"}]}),M=a(1),W=a(0),X=a(999),Y=a([]),Z=a({}),aa=async()=>{const a=await g({page:M.value,pageSize:X.value,...Z.value});0===a.code&&(Y.value=a.data.list,W.value=a.data.total,M.value=a.data.page,X.value=a.data.pageSize)};aa();const ea=(a,e)=>{K.value[a]=e},ta=a(null),la=a(null),oa=a(null),ra=(a,e)=>{const t=[ta,la,oa];e&&t[e].value.needConfirm&&(t[e].value.enterAndNext(),t[e].value.needConfirm=!1)},ia=a(null),ua=()=>{ia.value&&ia.value.resetFields(),J.value={authorityId:0,authorityName:"",parentId:0}},da=()=>{ua(),Q.value=!1,$.value=!1},sa=()=>{ia.value.validate((async a=>{if(a){switch(J.value.authorityId=Number(J.value.authorityId),G.value){case"add":0===(await k(J.value)).code&&(p({type:"success",message:"添加成功!"}),aa(),da());break;case"edit":0===(await _(J.value)).code&&(p({type:"success",message:"添加成功!"}),aa(),da());break;case"copy":{const a={authority:{authorityId:0,authorityName:"",datauthorityId:[],parentId:0},oldAuthorityId:0};a.authority.authorityId=J.value.authorityId,a.authority.authorityName=J.value.authorityName,a.authority.parentId=J.value.parentId,a.authority.dataAuthorityId=H.value.dataAuthorityId,a.oldAuthorityId=H.value.authorityId;0===(await w(a)).code&&(p({type:"success",message:"复制成功！"}),aa())}}ua(),Q.value=!1}}))},na=()=>{O.value=[{authorityId:0,authorityName:"根角色"}],ma(Y.value,O.value,!1)},ma=(a,e,t)=>{a&&a.forEach((a=>{if(a.children&&a.children.length){const l={authorityId:a.authorityId,authorityName:a.authorityName,disabled:t||a.authorityId===J.value.authorityId,children:[]};ma(a.children,l.children,t||a.authorityId===J.value.authorityId),e.push(l)}else{const l={authorityId:a.authorityId,authorityName:a.authorityName,disabled:t||a.authorityId===J.value.authorityId};e.push(l)}}))},pa=a=>{ua(),P.value="新增角色",G.value="add",J.value.parentId=a,na(),Q.value=!0};return(a,g)=>{const w=y,_=h,k=b,S=D,$=A,W=I,X=v,Z=E,ua=C,ma=B,ya=f,ha=c;return e(),t("div",q,[l(U,{title:"注：右上角头像下拉可切换角色"}),o("div",z,[o("div",F,[l(w,{type:"primary",icon:"plus",onClick:g[0]||(g[0]=a=>pa(0))},{default:r((()=>[i("新增角色")])),_:1}),l(_,{class:"cursor-pointer"},{default:r((()=>[l(u(d))])),_:1})]),l(S,{data:Y.value,"tree-props":{children:"children",hasChildren:"hasChildren"},"row-key":"authorityId",style:{width:"100%"}},{default:r((()=>[l(k,{label:"角色ID","min-width":"180",prop:"authorityId"}),l(k,{align:"left",label:"角色名称","min-width":"180",prop:"authorityName"}),l(k,{align:"left",label:"操作",width:"460"},{default:r((a=>[l(w,{icon:"setting",type:"primary",link:"",onClick:e=>{return t=a.row,T.value=!0,void(K.value=t);var t}},{default:r((()=>[i("设置权限")])),_:2},1032,["onClick"]),l(w,{icon:"plus",type:"primary",link:"",onClick:e=>pa(a.row.authorityId)},{default:r((()=>[i("新增子角色")])),_:2},1032,["onClick"]),l(w,{icon:"copy-document",type:"primary",link:"",onClick:e=>(a=>{na(),P.value="拷贝角色",G.value="copy";for(const e in J.value)J.value[e]=a[e];H.value=a,Q.value=!0})(a.row)},{default:r((()=>[i("拷贝")])),_:2},1032,["onClick"]),l(w,{icon:"edit",type:"primary",link:"",onClick:e=>(a=>{na(),P.value="编辑角色",G.value="edit";for(const e in J.value)J.value[e]=a[e];na(),Q.value=!0})(a.row)},{default:r((()=>[i("编辑")])),_:2},1032,["onClick"]),l(w,{icon:"delete",type:"primary",link:"",onClick:e=>{return t=a.row,void m.confirm("此操作将永久删除该角色, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((async()=>{0===(await j({authorityId:t.authorityId})).code&&(p({type:"success",message:"删除成功!"}),1===Y.value.length&&M.value>1&&M.value--,aa())})).catch((()=>{p({type:"info",message:"已取消删除"})}));var t}},{default:r((()=>[i("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])]),l(ua,{modelValue:Q.value,"onUpdate:modelValue":g[4]||(g[4]=a=>Q.value=a),title:P.value},{footer:r((()=>[o("div",R,[l(w,{onClick:da},{default:r((()=>[i("取 消")])),_:1}),l(w,{type:"primary",onClick:sa},{default:r((()=>[i("确 定")])),_:1})])])),default:r((()=>[l(Z,{ref_key:"authorityForm",ref:ia,model:J.value,rules:L.value,"label-width":"80px"},{default:r((()=>[l(W,{label:"父级角色",prop:"parentId"},{default:r((()=>[l($,{modelValue:J.value.parentId,"onUpdate:modelValue":g[1]||(g[1]=a=>J.value.parentId=a),style:{width:"100%"},disabled:"add"===G.value,options:O.value,props:{checkStrictly:!0,label:"authorityName",value:"authorityId",disabled:"disabled",emitPath:!1},"show-all-levels":!1,filterable:""},null,8,["modelValue","disabled","options"])])),_:1}),l(W,{label:"角色ID",prop:"authorityId"},{default:r((()=>[l(X,{modelValue:J.value.authorityId,"onUpdate:modelValue":g[2]||(g[2]=a=>J.value.authorityId=a),disabled:"edit"===G.value,autocomplete:"off",maxlength:"15"},null,8,["modelValue","disabled"])])),_:1}),l(W,{label:"角色姓名",prop:"authorityName"},{default:r((()=>[l(X,{modelValue:J.value.authorityName,"onUpdate:modelValue":g[3]||(g[3]=a=>J.value.authorityName=a),autocomplete:"off"},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue","title"]),T.value?(e(),s(ha,{key:0,modelValue:T.value,"onUpdate:modelValue":g[5]||(g[5]=a=>T.value=a),"with-header":!1,size:"40%",title:"角色配置"},{default:r((()=>[l(ya,{"before-leave":ra,type:"border-card"},{default:r((()=>[l(ma,{label:"角色菜单"},{default:r((()=>[l(N,{ref_key:"menus",ref:ta,row:K.value,onChangeRow:ea},null,8,["row"])])),_:1}),l(ma,{label:"角色api"},{default:r((()=>[l(V,{ref_key:"apis",ref:la,row:K.value,onChangeRow:ea},null,8,["row"])])),_:1}),l(ma,{label:"资源权限"},{default:r((()=>[l(x,{ref_key:"datas",ref:oa,authority:Y.value,row:K.value,onChangeRow:ea},null,8,["authority","row"])])),_:1})])),_:1})])),_:1},8,["modelValue"])):n("",!0)])}}});export{S as default};
