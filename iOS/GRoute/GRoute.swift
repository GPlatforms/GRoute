//
//  GRoute.swift
//  GRoute
//
//  Created by admin on 17/3/16.
//  Copyright © 2017年 admin. All rights reserved.
//

import Foundation
import Alamofire

enum GRouteResult {
    case Success()
    case Fail(Error?)
}

open class GRouteManager {
    public static let sharedInstance = GRouteManager()
    
    private init() {
        
    }
    
    public var urlConfig:[String] = []
    
    public var originDict:[String:Any] = [:]
    
    public var app_id = ""
    
    public var time = ""
    
    public var sign = ""
    
    public var urls:[String] = []
    
    public var isAvailable = false
    
    public func update(sucessCallback:@escaping () -> Void) {
        for item in urls {
            getRouteConfigFromServer(item, parameters: ["app_id":app_id,"timestamp":"\(time)","sign":sign], callback: { (res) in
                switch res {
                case .Success() :
                    self.isAvailable = true
                    sucessCallback()
                    return
                case .Fail(_):
                    break
                }
            })
        }
    }
    
    func getRouteConfigFromServer(_ url: String,
                                         method: HTTPMethod = .get,
                                         parameters: Parameters? = nil,
                                         encoding: ParameterEncoding = URLEncoding.default,
                                         headers: HTTPHeaders? = nil,
                                         callback:@escaping ((GRouteResult) -> Void))  {
        Alamofire.request(url, method: method, parameters: parameters, encoding: encoding, headers: headers).responseJSON { [weak self] (response:DataResponse<Any>) in
            debugPrint("Request: \(String(describing: response.request))")
            debugPrint("Response: \(String(describing: response.response))")
            debugPrint("Error: \(String(describing: response.error))")
            
            
            switch response.result {
            case .success(let value):
                if let jsonObj = value as? [String:Any] {
                    print("jsonObj:\(jsonObj)")
                    if jsonObj["code"] as? Int != 200 {
                        callback(GRouteResult.Fail(response.error))
                        break
                    }
                    self?.originDict = jsonObj
                    if let newurls = self?.originDict["base_url"] as? Array<String> {
                        self?.urlConfig = newurls
                    }
                }
                
                callback(GRouteResult.Success())
                break
            case .failure(_):
                callback(GRouteResult.Fail(response.error))
                break
            }
        }
    }
    
    
    
    public func get() -> String? {
        return urlConfig.first
    }
    
    public func get<T>(key : String) -> T? {
        return originDict[key] as? T
    }
    
    private func textMatch(text: String, pattern: String) -> Bool {
        return NSPredicate(format: "SELF MATCHES %@", pattern).evaluate(with: text)
    }
}
