(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["risks"],{4915:function(t,s,a){"use strict";var i=a("715e"),c=a.n(i);c.a},"715e":function(t,s,a){},"8e5b":function(t,s,a){"use strict";a.r(s);var i=function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("div",[t._m(0),t.isLoad?a("div",{staticClass:"container",staticStyle:{"padding-top":"30px"}},[a("div",{staticClass:"col-md-12 col-xs-12",staticStyle:{"border-bottom":"1px solid black","margint-top":"30px"}},[a("div",{staticClass:"row"},[t._m(1),a("div",{staticClass:"col-md-8 col-xs-8",staticStyle:{padding:"10px"}},[a("p",{staticClass:"sugestions"},[t._v(t._s(t.risk.description))])])])]),a("div",{staticClass:"col-md-12 col-xs-12",staticStyle:{"border-bottom":"1px solid black"}},[a("div",{staticClass:"row"},[t._m(2),a("div",{staticClass:"col-md-8 col-xs-8",staticStyle:{padding:"10px"}},[a("p",{staticClass:"sugestions"},[t._v(t._s(t.risk.stats))])])])]),t._m(3),a("div",{staticClass:"col-md-12 col-xs-12"},[a("div",{staticClass:"row"},[t._m(4),a("div",{staticClass:"col-md-8 col-xs-8",staticStyle:{padding:"10px"}},[a("p",{staticClass:"sugestions"},[t._v(t._s(t.risk.riskfactors))])])])])]):t._e(),a("possible-solution")],1)},c=[function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("div",{staticClass:"col-md-12 col-xs-12"},[a("div",{staticClass:"row"},[a("div",{staticClass:"page-title"},[a("h1",[t._v("Innefective fixes cost budge")]),a("div",{staticClass:"container",staticStyle:{"text-align":"end"}},[a("p",{attrs:{href:""}},[t._v("Accuracy: "),a("span",{staticStyle:{color:"#669999"}},[t._v("High")])]),a("p",{attrs:{href:""}},[t._v("Risk: "),a("span",{staticStyle:{color:"#669999"}},[t._v("medium")])])])])])])},function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("div",{staticClass:"col-md-3 col-xs-3",staticStyle:{"padding-bottom":"10px","margin-right":"10px"}},[a("div",{staticClass:"col-md-2"},[a("span",[a("i",{staticClass:"orange-icon mt-20 far fa-5x fa-file-alt"})]),a("h1",{staticClass:"risks-h1"},[t._v("Description")])])])},function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("div",{staticClass:"col-md-3 col-xs-3",staticStyle:{"padding-bottom":"10px","margin-right":"10px"}},[a("div",{staticClass:"col-md-2"},[a("span",[a("i",{staticClass:"orange-icon mt-20 fas fa-5x fa-chart-line"})]),a("h1",{staticClass:"risks-h1"},[t._v("Stats")])])])},function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("div",{staticClass:"col-md-12 col-xs-12",staticStyle:{"border-bottom":"1px solid black"}},[a("div",{staticClass:"row pt-20"},[a("div",{staticClass:"col-md-3 col-xs-3",staticStyle:{"padding-bottom":"10px","margin-right":"10px"}},[a("span",[a("i",{staticClass:"orange-icon mt-20 fas fa-5x fa-map-marker-alt"})]),a("h1",{staticClass:"risks-h1"},[t._v("Usual Pain\n            "),a("br"),t._v("Points")])]),a("div",{staticClass:"col-md-8 col-xs-8",staticStyle:{"padding-bottom":"10px"}},[a("div",[a("div",{staticClass:"col-md-3"},[a("div",{staticClass:"blocks medium"},[t._v("\n                Decreasing\n                profits\n              ")])]),a("div",{staticClass:"col-md-3"},[a("div",{staticClass:"blocks high"},[t._v("\n                Decreasing\n                profits\n              ")])]),a("div",{staticClass:"col-md-3"},[a("div",{staticClass:"blocks low"},[t._v("\n                Decreasing\n                profits\n              ")])]),a("div",{staticClass:"col-md-3"},[a("div",{staticClass:"blocks low"},[t._v("\n                Decreasing\n                profits\n              ")])])])])])])},function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("div",{staticClass:"col-md-3 col-xs-3",staticStyle:{"padding-bottom":"10px","margin-right":"10px"}},[a("h1",{staticClass:"risks-h1"},[t._v("Risks Factors")])])}],o=a("c93e"),e=(a("cadf"),a("551c"),a("097d"),a("bc3a")),l=a.n(e),n=a("b746"),r=a("8252"),d=a("9a4c"),v={name:"Risks",mixins:[n["a"]],components:{PossibleSolution:r["default"],BModalFullScreen:d["a"]},data:function(){return{risk:{},isLoad:!1,issueId:""}},mounted:function(){var t=this;this.issueId=window.sessionStorage.getItem("issueId"),l.a.get("".concat("/api","/issue/").concat(this.issueId)).then(function(s){if(!s.data.status)throw new Error("".concat(s.data.message));t.isLoad=!0,t.risk=Object(o["a"])({},s.data.data)}).catch(function(s){return t.errorMsg("Risk","".concat(s))})}},m=v,u=(a("c5aa"),a("2877")),p=Object(u["a"])(m,i,c,!1,null,"92f35e7e",null);p.options.__file="Risk.vue";s["default"]=p.exports},"9a4c":function(t,s,a){"use strict";var i=function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("div",{staticClass:"overlay",style:t.modalControl?"width:100%":"width:0",attrs:{id:"myNav"}},[a("span",{staticClass:"closebtn",on:{click:t.closeModal}},[a("i",{staticClass:"fas fa-times-circle"})]),a("div",{staticClass:"overlay-content"},[t._t("content")],2)])},c=[],o={name:"BModalFullScreen",props:{show:{type:Boolean,default:!1}},computed:{modalControl:{get:function(){return this.show},set:function(t){this.$emit("update:show",t)}}},methods:{closeModal:function(){this.$emit("close")}}},e=o,l=(a("4915"),a("2877")),n=Object(l["a"])(e,i,c,!1,null,"51718183",null);n.options.__file="BModalFullScreen.vue";s["a"]=n.exports},b3f9:function(t,s,a){},c5aa:function(t,s,a){"use strict";var i=a("b3f9"),c=a.n(i);c.a}}]);
//# sourceMappingURL=risks.fcb8d6cf.js.map