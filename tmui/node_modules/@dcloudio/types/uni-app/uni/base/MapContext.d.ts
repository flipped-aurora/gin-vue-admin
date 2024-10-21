import { AsyncOptions } from './Utils';

declare global {
  namespace UniNamespace {
    interface MapContext {
      /**
       * 获取当前地图中心的经纬度，返回的是 gcj02 坐标系，可以用于 uni.openLocation
       */
      getCenterLocation(options: MapContextGetCenterLocationOptions): void;
      /**
       * 将地图中心移动到当前定位点，需要配合map组件的show-location使用
       */
      moveToLocation(options: MapContextMoveToLocationOptions): void;
      /**
       * 平移marker，带动画
       */
      translateMarker(options: MapContextTranslateMarkerOptions): void;
      /**
       * 缩放视野展示所有经纬度
       */
      includePoints(options: MapContextIncludePointsOptions): void;
      /**
       * 获取当前地图的视野范围
       */
      getRegion(options: MapContextGetRegionOptions): void;
      /**
       * 获取当前地图的缩放级别
       */
      getScale(options: MapContextGetScaleOptions): void;
      /**
       * 添加个性化图层
       */
      addCustomLayer?: (options: MapContextAddCustomLayerOptions) => void;
      /**
       * 创建自定义图片图层，图片会随着地图缩放而缩放
       */
      addGroundOverlay(options: MapContextAddGroundOverlayOptions): void;
      /**
       * 添加 marker
       */
      addMarkers(options: MapContextAddMarkersOptions): void;
      /**
       * 获取屏幕上的点对应的经纬度，坐标原点为地图左上角
       */
      fromScreenLocation?: (options: MapContextFromScreenLocationOptions) => void;
      /**
       * 初始化点聚合的配置，未调用时采用默认配置
       */
      initMarkerCluster(options: MapContextInitMarkerClusterOptions): void;
      /**
       * 沿指定路径移动 marker，用于轨迹回放等场景。动画完成时触发回调事件，若动画进行中，对同一 marker 再次调用 moveAlong 方法，前一次的动画将被打断。
       */
      moveAlong(options: MapContextMoveAlongOptions): void;
      /**
       * 拉起地图APP选择导航。
       */
      openMapApp(options: MapContextOpenMapAppOptions): void;
      /**
       * 移除个性化图层
       */
      removeCustomLayer(options: MapContextRemoveCustomLayerOptions): void;
      /**
       * 移除自定义图片图层
       */
      removeGroundOverlay(options: MapContextRemoveGroundOverlayOptions): void;
      /**
       * 移除 marker
       */
      removeMarkers(options: MapContextRemoveMarkersOptions): void;
      /**
       * 设置地图中心点偏移，向后向下为增长，屏幕比例范围(0.25~0.75)，默认偏移为[0.5, 0.5]
       */
      setCenterOffset?: (options: MapContextSetCenterOffsetOptions) => void;
      /**
       * 获取经纬度对应的屏幕坐标，坐标原点为地图左上角。
       */
      toScreenLocation?: (options: MapContextToScreenLocationOptions) => void;
      /**
       * 更新自定义图片图层。
       */
      updateGroundOverlay(options: MapContextUpdateGroundOverlayOptions): void;
      /**
       * 监听地图事件。
       */
      on(event: 'markerClusterCreate' | 'markerClusterClick', callback: (...args: any[]) => any): void;
      /**
       * 获取原生地图对象 plus.maps.Map
       */
      $getAppMap(): any;
    }

    interface MapContextGetCenterLocationOptions extends AsyncOptions {
      /**
       * 接口调用成功的回调函数 ，res = { longitude: "经度", latitude: "纬度"}
       */
      success?: (result: LocationObject) => void;
    }

    interface MapContextMoveToLocationOptions extends AsyncOptions {
      /**
       * 纬度，浮点数，范围为-90~90，负数表示南纬
       */
      latitude?: number;
      /**
       * 经度，范围为-180~180，负数表示西经
       */
      longitude?: number;
    }

    interface LocationObject {
      /**
       * 纬度，浮点数，范围为-90~90，负数表示南纬
       */
      latitude: number;
      /**
       * 经度，范围为-180~180，负数表示西经
       */
      longitude: number;
    }

    interface MapContextTranslateMarkerOptions extends AsyncOptions {
      /**
       * 指定marker
       */
      markerId: number;
      /**
       * 指定marker移动到的目标点
       */
      destination: LocationObject;
      /**
       * 移动过程中是否自动旋转marker
       */
      autoRotate: boolean;
      /**
       * marker的旋转角度
       */
      rotate: number;
      /**
       * 平移和旋转同时进行，默认值false（仅微信小程序2.13.0支持）
       */
      moveWithRotate?: boolean;
      /**
       * 动画持续时长，默认值1000ms，平移与旋转分别计算
       */
      duration?: number;
      /**
       * 动画结束回调函数
       */
      animationEnd?: (result: any) => void;
    }

    interface MapContextIncludePointsOptions extends AsyncOptions {
      /**
       * 要显示在可视区域内的坐标点列表，[{latitude, longitude}]
       */
      points: LocationObject[];
      /**
       * 坐标点形成的矩形边缘到地图边缘的距离，单位像素。格式为[上,右,下,左]，安卓上只能识别数组第一项，上下左右的padding一致。开发者工具暂不支持padding参数。
       */
      padding?: number[];
    }

    interface MapContextGetRegionOptions extends AsyncOptions {
      /**
       * 接口调用成功的回调函数，res = {southwest, northeast}，西南角与东北角的经纬度
       */
      success?: (result: MapContextGetRegionResult) => void;
    }

    interface MapContextGetRegionResult {
      /**
       * 西南角的经纬度
       */
      southwest: LocationObject;
      /**
       * 东北角的经纬度
       */
      northeast: LocationObject;
    }

    interface MapContextGetScaleOptions extends AsyncOptions {
      /**
       * 接口调用成功的回调函数，res = {scale}
       */
      success?: (result: MapContextGetScaleResult) => void;
    }

    interface MapContextGetScaleResult {
      /**
       * 地图缩放级别
       */
      scale: number;
    }

    interface MapContextAddCustomLayerOptions extends AsyncOptions {
      /**
       * 个性化图层id
       */
      layerId: string;
    }

    interface MapContextAddGroundOverlayOptions extends AsyncOptions {
      /**
       * 图片图层 id
       */
      id: string;
      /**
       * 图片路径，支持网络图片、临时路径、代码包路径
       */
      src: string;
      /**
       * 图片覆盖的经纬度范围
       */
      bounds: Bounds;
      /**
       * 是否可见
       */
      visible?: boolean;
      /**
       * 图层绘制顺序
       */
      zIndex?: number;
      /**
       * 图层透明度
       */
      opacity?: number;
    }

    interface Bounds {
      /**
       * 西南角的经纬度
       */
      southwest: LocationObject;
      /**
       * 东北角的经纬度
       */
      northeast: LocationObject;
    }

    interface MapContextAddMarkersOptions extends AsyncOptions {
      /**
       * 同传入 map 组件的 marker 属性
       */
      markers: any[];
      /**
       * 是否先清空地图上所有 marker
       */
      clear: boolean;
    }

    interface MapContextFromScreenLocationOptions extends AsyncOptions {
      /**
       * x 坐标值
       */
      x: number;
      /**
       * y 坐标值
       */
      y: number;
      /**
       * 接口调用成功的回调函数
       */
      success?: (result: LocationObject) => void;
    }

    interface MapContextInitMarkerClusterOptions extends AsyncOptions {
      /**
       * 启用默认的聚合样式
       */
      enableDefaultStyle: boolean;
      /**
       * 点击已经聚合的标记点时是否实现聚合分离
       */
      zoomOnClick: boolean;
      /**
       * 聚合算法的可聚合距离，即距离小于该值的点会聚合至一起，以像素为单位
       */
      gridSize: number;
    }

    interface MapContextMoveAlongOptions extends AsyncOptions {
      /**
       * 指定 marker
       */
      markerId: number;
      /**
       * 移动路径的坐标串，坐标点格式 {longitude, latitude}
       */
      path: LocationObject[];
      /**
       * 根据路径方向自动改变 marker 的旋转角度
       */
      autoRotate?: boolean;
      /**
       * 平滑移动的时间
       */
      duration: number;
    }

    interface MapContextOpenMapAppOptions extends AsyncOptions {
      /**
       * 目的地名称
       */
      destination: string;
      /**
       * 目的地纬度
       */
      latitude: number;
      /**
       * 目的地经度
       */
      longitude: number;
    }

    interface MapContextRemoveCustomLayerOptions extends AsyncOptions {
      /**
       * 个性化图层id
       */
      layerId: string;
    }

    interface MapContextRemoveGroundOverlayOptions extends AsyncOptions {
      /**
       * 图片图层 id
       */
      id: string;
    }

    interface MapContextRemoveMarkersOptions extends AsyncOptions {
      /**
       * 要被删除的marker的id属性组成的数组
       */
      markerIds: any[];
    }

    interface MapContextSetCenterOffsetOptions {
      /**
       * 偏移量，两位数组
       */
      offset: number[];
      /**
       * 接口调用成功的回调函数
       */
      success?: (result: any) => void;
      /**
       * 接口调用失败的回调函数
       */
      fail?: (result: any) => void;
      /**
       * 接口调用结束的回调函数（调用成功、失败都会执行）
       */
      complete?: (result: any) => void;
    }

    interface MapContextToScreenLocationOptions extends AsyncOptions {
      /**
       * 纬度
       */
      latitude: number;
      /**
       * 经度
       */
      longitude: number;
    }

    interface MapContextUpdateGroundOverlayOptions extends AsyncOptions {
      /**
       * 图片图层 id
       */
      id: string;
      /**
       * 图片路径，支持网络图片、临时路径、代码包路径
       */
      src: string;
      /**
       * 图片路径，支持网络图片、临时路径、代码包路径
       */
      bounds: Bounds;
      /**
       * 是否可见
       */
      visible?: boolean;
      /**
       * 图层绘制顺序
       */
      zIndex?: number;
      /**
       * 图层透明度
       */
      opacity?: number;
    }
  }

  interface Uni {
    /**
     * 创建并返回 map 上下文 mapContext 对象
     *
     * 文档: [http://uniapp.dcloud.io/api/location/map?id=createmapcontext](http://uniapp.dcloud.io/api/location/map?id=createmapcontext)
     */
    createMapContext(mapId: string, currentComponent?: any): UniNamespace.MapContext;
  }
}
