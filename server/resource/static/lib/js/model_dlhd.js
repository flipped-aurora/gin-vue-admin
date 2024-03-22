// JavaScript Document
var dlhd_dgg = {
	 percent:0,
	 timers:"",
	 init:function(e){
		 /*wow*/
		  var wow = new WOW({
			boxClass: 'wow',
			animateClass: 'animated',
			offset: 0,
			mobile: true,
			live: true
		});
		wow.init();
		dlhd_dgg.resize();
		$(window).resize(function(){
			 dlhd_dgg.resize();
		});
		
		/*scroll*/
		$(window).scroll(function(){
			t = $(document).scrollTop();
			if(t > 50){
				$(".header").addClass("fixedmenu");
				$(".phone_nav").addClass("fixedmenu");
			}else{
				$('.header').removeClass("fixedmenu");
				$(".phone_nav").removeClass("fixedmenu");
			}   
			if(t > 100){
			    $(".back_top").show();
			}else
			{
				$(".back_top").hide();
			}
			    
		   });  	
		/*返回顶部*/	
       $('.back_top').click(function(){ 
			$('body,html').animate({
				scrollTop: 0
			},
			800);//点击回到顶部按钮，缓懂回到顶部,数字越小越快
			return false; 
		}); 
		//搜索
		 $(".i_bcss .i_bcss_t").click(function(){
		 $(this).parent().find(".i_bcss_c").toggle();
	    });
	   $(".i_ssbn span").click(function(){
		   $(this).parent().find(".i_ssbn_div").toggle();
		});
		/*pc导航*/
		$(".i_logo_ul .nav_menu li").hover(function(){
		   $(this).find(".a_nav").addClass("active");
		   $(this).find("dl").stop(false,true).slideDown(500);
		},function(){
			$(this).find(".a_nav").removeClass("active");
			$(this).find("dl").stop(false,true).hide();
		}); 
		
		$(".i_flogo_r .s2").click(function(){
			$(this).parent().find(".link_text").toggle();
		});
		
		/*pc导航*/
		
		/*手机导航*/
		var btn = true;
		$(".nav_but_box").click(function() {
			$(this).parent().toggleClass("phone_nav_one")
			$("body").toggleClass('hided');
			if (btn) {
				$(".phone_nav .nav_but").css("background", "none");
				btn = false;
			} else {
				$(".phone_nav .nav_but").css("background", "#1f3c8a");
				btn = true;
			}
		})
		$(".nav_main b").click(function() {
			$(this).parent().toggleClass("b_one").siblings().removeClass("b_one");
		});
		$(".nav_main li .ul b").click(function() {
			$(this).parent().toggleClass("b_one1").siblings().removeClass("b_one1");
		});
		 $(".phone_nav .nav_main ul li").each(function(i){
		 
		  $(this).find(".ul li:last-child").css({"border":"0px"});
	    
	   });
		/*手机导航*/
		
		$(".a_zxlmenu_c li b").click(function(){
			 
			$(".a_zxlmenu_c li .aone").removeClass("activetwo");
			$(".a_zxlmenu_c li dl").slideUp();
		     var dlshow=$(this).parent().find("dl").length;
		     var dis=$(this).parent().find("dl").css("display");
			   if(dlshow>0&&dis=="none")
			   {
			   $(this).parent().find(".aone").addClass("activetwo");
			  $(this).parent().find("dl").stop(true,false).slideDown(200);
			   }	  
			 
			 
			});
		 
	
			$(".a_zxlmenu .a_zxlmenu_t").click(function(){
			$(this).parent().find(".a_zxlmenu_c").toggle(500);
			});
		
	},
	 loading:function()
	 {
		 /*loading1*/
		setTimeout(function(){$(".loading").fadeOut();},1000);
		 /*loading2*/
		  dlhd_dgg.timers = setInterval(function() {
			dlhd_dgg.percent++;
			if (dlhd_dgg.percent > 99) {
				dlhd_dgg.index();
				return;
			}
			$(".loadcircle .baifen").html(dlhd_dgg.percent + " %");
		}, 30)
		/*loading2*/
		/*loading3*/
		 setTimeout(function(){$(".layer").addClass("show");},1000);
	     setTimeout(function(){$(".filler-left,.filler-right").css({height:'96%'});},2500);
	     setTimeout(function(){$(".filler-top,.filler-bottom").css({width:'98%'});},1500);
		/*loading3*/
		
		
	 },
	 index:function()
	 {
		 $(window).resize(function(){
			dlhd_dgg.resize();
			});
		 /*loading2*/
		 $(".loadcircle .baifen").html("100%")
		 clearInterval(dlhd_dgg.timers);
		 $(".loadBox").addClass("active");
		 /*loading2*/
		
		 
		 new Swiper('.banner-swiper', {
			   pagination: {
				el: '.banner-pagination',
				dynamicBullets: false,
				clickable: true,
			},
			  fadeEffect: {
				crossFade: true,
			  },
			autoplay: true,
			speed: 1000,
			loop: true,
			 navigation: {
			  nextEl: ".banner_next",
			  prevEl: ".banner_prev",
			},
            });
	    new Swiper('.hd-swiper', {
			pagination: {
				el: '.hd-pagination',
				dynamicBullets: false,
				clickable :true
			  },
			autoplay:true,
			slidesPerView: 1,
			speed:1000,
			loop:true,
			navigation: {
            nextEl: '.hd_next',
            prevEl: '.hd_prev'
            },
			
		});
		
		$(".frame_content_box1_r_t a").eq(0).addClass("activetm");
		$(".frame_content_box1_r_c .frame_content_box1_r_con").eq(0).show();
		$(".frame_content_box1_r_t a").hover(function(){
			var num=$(this).index();
			$(this).addClass("activetm").siblings().removeClass("activetm");
			$(".frame_content_box1_r_c .frame_content_box1_r_con").hide().eq(num).show();
			
		});
		
		new Swiper('.project_swiper', {
		 pagination: {
			el: '.project_pagination',
			dynamicBullets: false,
			clickable: true
		  },
		  navigation: {
			nextEl:'.project_next',
			prevEl:'.project_prev'
			},
			 autoplay:true,
			 speed:1000,
			 loop:false,
			 slidesPerView:3,
			 spaceBetween :30,
			 breakpoints: { 
			640: {
			  slidesPerView: 1,
			  spaceBetween : 10,
			},
			768: {
			  slidesPerView: 2,
			   spaceBetween : 10,
			},
			//当宽度小于等于992
			1200: {
			  slidesPerView: 3
			 
			}
			
			}
	   });
			
	  
		 
	 },map:function(a,b,c)
	 {
		 var map = new BMap.Map(a);
			map.addControl(new BMap.NavigationControl());
			map.addControl(new BMap.ScaleControl());    
			map.addControl(new BMap.OverviewMapControl()); 
			map.enableScrollWheelZoom(true);
			map.centerAndZoom(new BMap.Point(b,c), 15);
			var marker1 = new BMap.Marker(new BMap.Point(b,c));  // 创建标注
			map.addOverlay(marker1);              // 将标注添加到地图中
			marker1.setAnimation(BMAP_ANIMATION_BOUNCE); //跳动的动画
	 },resize:function()
	 {
		 var bgWidth=document.documentElement.clientWidth;
		 var bgHeight=document.documentElement.clientHeight;
		 /*setTimeout(function(){
		 $("#banner,.home-swiper").css({"height":bgHeight-90+"px"});
		 },200);	*/
		 
		 
	 },resources:function()
	 {
		 
	  $(".a_careers_c dd").eq(0).find(".div").addClass("activeds");
		$(".a_careers_c dd").eq(0).find(".div_l").slideDown(500);
		$(".a_careers_c dd").click(function(){
			$(".a_careers_c dd .div").removeClass("activeds");
			$(".a_careers_c dd .div_l").slideUp(500);
			var dis=$(this).find(".div_l").css("display");
			var dlshow=$(this).find(".div_l").length;
			 if(dlshow>0&&dis=="none")
			{
				$(this).find(".div").addClass("activeds");
				$(this).find(".div_l").slideDown(500);
			}
	   });
		
		
	   
	 },products:function(){
		
		
	  
		var swiper1= new Swiper(".product_swiper_left", {
		direction: 'vertical',
		slidesPerView:4,
		watchSlidesProgress: true,
        navigation: {
          nextEl: ".product_left_next",
          prevEl: ".product_left_prev",
        },
		 breakpoints: { 
			640: {
			},
			768: {
				
			  
			},
		 }
      });
	  var swiper2= new Swiper(".product_swiper_right", {
        navigation: {
          nextEl: ".product_right_next",
          prevEl: ".product_right_prev",
        },
        thumbs: {
          swiper: swiper1,
        },
		 breakpoints: { 
			640: {
			  thumbs: {
			  swiper:'',
			},
			},
			768: {
				
			  
			},
		 }
      });
		
	}
	 
}
dlhd_dgg.init();