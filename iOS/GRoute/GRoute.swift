//
//  GRoute.swift
//  GRoute
//
//  Created by weibo on 17/3/16.
//  Copyright © 2017年 lez. All rights reserved.
//

import Foundation
import Alamofire

public enum GRouteResult {
    case Success()
    case Fail(Error?)
}

public class Rule {
    var reg = ""
    var url = ""
    
    init() {
        
    }
    
    init(newReg:String?,newURL:String?) {
        reg = newReg ?? ""
        url = newURL ?? ""
    }
}



public class GRouteManager {
    public static let sharedInstance = GRouteManager()
    
    private init() {
        
    }
    
    public var urlConfig:[Rule] = []
    
    public var originDict:[String:Any] = [:]
    
    public func getRouteConfigFromServer(_ url: String,
                                         method: HTTPMethod = .get,
                                         parameters: Parameters? = nil,
                                         encoding: ParameterEncoding = URLEncoding.default,
                                         headers: HTTPHeaders? = nil,
                                         callback:@escaping ((GRouteResult) -> Void))  {
        Alamofire.request(url).responseJSON { (response:DataResponse<Any>) in
            debugPrint("Request: \(response.request)")
            debugPrint("Response: \(response.response)")
            debugPrint("Error: \(response.error)")
            
            
            switch response.result {
            case .success(let value):
                if let data = value as? NSData {
                    debugPrint("data")
                }
                if let str = value as? String {
                    debugPrint("String")
                }
                /*if let jsonObj = (try? JSONSerialization.jsonObject(with: value, options: .allowFragments) ) as? [AnyHashable:Any] {
                    
                }*/
                
                callback(GRouteResult.Success())
                break
            case .failure(_):
                callback(GRouteResult.Fail(response.error))
                break
            }
        }
    }
    
    
    
    public func getBaseUrl(functionName : String) -> String? {
        for item in self.urlConfig {
            if textMatch(text: functionName, pattern: item.reg) {
                return item.url
            }
        }
        return nil
    }
    
    public func get<T>(key : String) -> T? {
        return originDict[key] as? T
    }
    
    private func textMatch(text: String, pattern: String) -> Bool {
        return NSPredicate(format: "SELF MATCHES %@", pattern).evaluate(with: text)
    }
}
