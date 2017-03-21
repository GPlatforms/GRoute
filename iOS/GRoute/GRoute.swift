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
    case Success([Rule])
    case Fail(Error?)
}

class Rule {
    var reg = ""
    var url = ""
    
    init() {
        
    }
    
    init(newReg:String?,newURL:String?) {
        reg = newReg ?? ""
        url = newURL ?? ""
    }
}

class GRouteClient {
    static let sharedInstance = GRouteClient()
    
    init() {
        
    }
    
    public var routeConfig:[Rule] = []
    
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
                let res:[Rule] = []
                callback(GRouteResult.Success(res))
                break
            case .failure(_):
                callback(GRouteResult.Fail(response.error))
                break
            }
        }
    }
    
    
    
    public func match(functionName:String) -> String? {
        for item in self.routeConfig {
            if textMatch(text: functionName, pattern: item.reg) {
                return item.url
            }
        }
        return nil
    }
    
    public func textMatch(text: String, pattern: String) -> Bool {
        debugPrint("text: \(text)")
        debugPrint("pattern: \(pattern)")
        return NSPredicate(format: "SELF MATCHES %@", pattern).evaluate(with: text)
    }
}
