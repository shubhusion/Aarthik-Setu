import 'package:aarthik_setu/pages/mobile/auth/components/phone_number_form.dart';
import 'package:aarthik_setu/pages/mobile/auth/components/sign_in_options.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:responsive_framework/responsive_framework.dart';
import '../../../constants/app_constants.dart';
import '../../../cubit/phone_form_cubit.dart';

class SignInMobile extends StatelessWidget {
  const SignInMobile({super.key});

  @override
  Widget build(BuildContext context) {
    return ResponsiveScaledBox(
      width: AppConstants.mobileScaleWidth,
      child: Scaffold(
        body: SizedBox.expand(
          child: SingleChildScrollView(
            child: Column(
              children: [
                const SizedBox(height: 200),
                Text(
                  'Aarthik Setu',
                  style: GoogleFonts.poppins(fontSize: 55),
                ),
                const SizedBox(height: 80),
                BlocBuilder<PhoneFormCubit, PhoneFormState>(
                  builder: (context, state) {
                    return Container(
                      width: 350,
                      decoration: BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.circular(20),
                        boxShadow: [
                          BoxShadow(
                            color: Colors.grey.withOpacity(0.5),
                            spreadRadius: 5,
                            blurRadius: 7,
                            offset: const Offset(0, 3),
                          ),
                        ],
                      ),
                      child: IntrinsicHeight(
                        child: Padding(padding: const EdgeInsets.only(top:20, bottom: 40),
                        child: (state as PhoneForm).isPhoneInputOpen
                            ? const PhoneNumberFormMobile()
                            : const SignInOptionsMobile(),
                        ),
                      )
                    );
                  },
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
