//
//  test1.swift
//  GRoute
//
//  Created by admin on 2017/9/21.
//  Copyright © 2017年 lez. All rights reserved.
//

import XCTest
@testable import GRoute

class test1: XCTestCase {
    
    override func setUp() {
        super.setUp()
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }
    
    override func tearDown() {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
        super.tearDown()
    }
    
    func testPredicate() {
        let ex = expectation(description: "1")
        
        let gRoute = GRouteManager.sharedInstance
        print("\(Int(NSDate().timeIntervalSince1970))")
        let time = "1505977676"//"\(Int(NSDate().timeIntervalSince1970))"
        gRoute.getConfig(app_id: "11235", time: time, sign: "3f9a612fac3014f80794b74392917e8113b7a052", urls: ["http://wenzb.com/groute/v1/config","http://121.40.106.138/groute/v1/config"]) {
            print("baseURL:",gRoute.getBaseUrl())
            print(gRoute.originDict)
            ex.fulfill()
        }
        //
        waitForExpectations(timeout: 100) { (err) in
            print(err)
        }
        // This is an example of a functional test case.
        // Use XCTAssert and related functions to verify your tests produce the correct results.
    }
    
    func testExample() {
        // This is an example of a functional test case.
        // Use XCTAssert and related functions to verify your tests produce the correct results.
    }
    
    func testPerformanceExample() {
        // This is an example of a performance test case.
        self.measure {
            // Put the code you want to measure the time of here.
        }
    }
    
}
