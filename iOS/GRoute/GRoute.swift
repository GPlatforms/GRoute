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
        Alamofire.request(url).responseJSON { [weak self] (response:DataResponse<Any>) in
            debugPrint("Request: \(response.request)")
            debugPrint("Response: \(response.response)")
            debugPrint("Error: \(response.error)")
            
            
            switch response.result {
            case .success(let value):
                if let jsonObj = value as? [AnyHashable:Any] {
                    if jsonObj["code"] as? Int != 200 {
                        callback(GRouteResult.Fail(response.error))
                        break
                    }
                    guard let data = jsonObj["data"] as? [String:Any] else {
                        callback(GRouteResult.Fail(response.error))
                        break
                    }
                    self?.originDict = data
                    if let urlConfig = self?.originDict["base_url"] as? Array<[String:Any]> {
                        self?.urlConfig = urlConfig.map({ (obj) -> Rule in
                            return Rule.init(newReg: obj["reg"] as? String, newURL: obj["url"] as? String)
                        })
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
    
    
    
    public func getBaseUrl(functionName : String = "*") -> String? {
        for item in self.urlConfig {
            if item.reg == functionName {
                return item.url
            }
        }
        //return urlConfig
        let all = urlConfig.filter({ (rule) -> Bool in
            return rule.reg == "*"
        })
        return all.first?.url
    }
    
    public func get<T>(key : String) -> T? {
        return originDict[key] as? T
    }
    
    private func textMatch(text: String, pattern: String) -> Bool {
        return NSPredicate(format: "SELF MATCHES %@", pattern).evaluate(with: text)
    }
}
