//
//  GRoute.swift
//  GRoute
//
//  Created by weibo on 17/3/16.
//  Copyright © 2017年 lez. All rights reserved.
//

import Foundation
import Alamofire

enum GRouteResult {
    case Success([String:Any])
    case Fail(Error?)
}

class GRoute {
    static let sharedInstance = GRoute()
    
    init() {
        print(GRoute.sharedInstance.match(text: "adsfa!", pattern: ".*"))
    }
    
    private(set) var routeConfig:[String:Any] = [:]
    
    public func getRouteConfigFromServer(_ url: String,
                                         method: HTTPMethod = .get,
                                         parameters: Parameters? = nil,
                                         encoding: ParameterEncoding = URLEncoding.default,
                                         headers: HTTPHeaders? = nil,
                                         callback:@escaping ((GRouteResult) -> Void))  {
        Alamofire.request(url).responseJSON { response in
            debugPrint("Request: \(response.request)")
            debugPrint("Response: \(response.response)")
            debugPrint("Error: \(response.error)")

            
            if let json = response.result.value {
                debugPrint("JSON: \(json)")
            }
            switch response.result {
            case .success(_):
                var res:[String:Any] = [:]
                callback(GRouteResult.Success(res))
                break
            case .failure(_):
                callback(GRouteResult.Fail(response.error))
                break
            }
        }
    }
    
    
    
    public func match(moduleName:String) -> String {
        for (key,value) in self.routeConfig {
            if match(text: moduleName, pattern: key) {
                return ""
            }
        }
        return ""
    }
    
    public func match(text: String, pattern: String) -> Bool {
        return NSPredicate(format: "SELF MATCHES %@", pattern).evaluate(with: text)
    }
}

let GRouteClient = GRoute.sharedInstance
